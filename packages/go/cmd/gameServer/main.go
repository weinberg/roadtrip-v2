package main

import (
	"context"
	"errors"
	"fmt"
	. "github.com/brickshot/roadtrip-v2/internal/gameServer"
	g "github.com/brickshot/roadtrip-v2/internal/gameServer/grpc"
	db "github.com/brickshot/roadtrip-v2/internal/prisma"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var dbHost = os.Getenv("DATABASE_HOST")
var dbPort = os.Getenv("DATABASE_PORT")
var dbUser = os.Getenv("DATABASE_USER")
var dbPass = os.Getenv("DATABASE_PASS")
var dbName = os.Getenv("DATABASE_NAME")
var dbURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?schema=public",
	dbHost, dbPort, dbUser, dbPass, dbName)

const port = "9066"
const updateInterval = time.Second

type gameServer struct {
	g.UnimplementedRoadTripGameServer
}

/******************************
 *
 * Data
 *
 ******************************/

type gameData struct {
	cars map[string]Car
	routes map[string]Route
}

var data gameData = gameData{cars: make(map[string]Car), routes: make(map[string]Route)}
var dataMutex sync.RWMutex // guards data

var lastTick time.Time
var client *db.PrismaClient

/******************************
 *
 * Main
 *
 ******************************/

func init() {
	lastTick = time.Now()
}

func main() {
	fmt.Printf("GameServer starting...\n")

	// Prisma
	setupPrisma()
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	// load data from db
	<-loadRoutes()
	<-loadCars()

	// Start server
	c := make(chan int)
	go mainLoop(c)
	go StartServer(c)
	<-c
}

func setupPrisma() {
	// Setup prisma
	client = db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
}

func loadRoutes() <-chan int32 {
	fmt.Printf("Loading routes...\n")

	dataMutex.Lock()
	defer dataMutex.Unlock()

	c := make(chan int32)
	go func() {
		defer close(c)
		ctx := context.Background()
		rs, _ := client.Route.FindMany().With(
			db.Route.Ways.Fetch().With(
				db.WaysOnRoutes.Way.Fetch().With(
					db.Way.Nodes.Fetch().With(
						db.NodesOnWays.Node.Fetch(),
					).OrderBy(
						db.NodesOnWays.Sequence.Order(db.SortOrderAsc),
					),
				),
			).OrderBy(
				db.WaysOnRoutes.Sequence.Order(db.SortOrderAsc),
			),
		).Exec(ctx)
		routes := unmarshallRoutes(rs)

		for _, route := range routes {
			data.routes[route.Id] = route
		}

		c <- 1
	}()
	return c
}

func unmarshallRoutes(rs []db.RouteModel) []Route {
	res := []Route{}
	for _, r := range rs {
		route := Route{
			Id:    r.ID,
			Nodes: []Node{},
		}
		for _, w := range r.Ways() {
			for _, n := range w.RelationsWaysOnRoutes.Way.Nodes() {
				route.Nodes = append(route.Nodes, Node{
					Miles: n.RelationsNodesOnWays.Node.Miles,
					Id:    n.RelationsNodesOnWays.Node.ID,
				})
			}
		}
		res = append(res, route)
	}

	return res
}

func loadCars() <-chan int32 {
	fmt.Printf("Loading cars...\n")

	dataMutex.Lock()
	defer dataMutex.Unlock()

	c := make(chan int32)
	go func() {
		defer close(c)
		ctx := context.Background()
		cs, _ := client.Car.FindMany().Exec(ctx)
		cars := unmarshallCars(cs)

		for _, car := range cars {
			data.cars[car.Id] = car
		}

		c <- 1
	}()
	return c
}

func unmarshallCars(cs []db.CarModel) []Car {
	res := []Car{}
	for _, c := range cs {
		routeIndex, ok := c.RouteIndex()
		if !ok {
			continue
		}
		nodeMiles, ok := c.NodeMiles()
		if !ok {
			continue
		}
		routeId, ok := c.RouteID()
		if !ok {
			continue
		}
		mph, ok := c.Mph()
		if !ok {
			continue
		}
		car := Car{
			Id:         c.ID,
			RouteId:    routeId,
			RouteIndex: int32(routeIndex),
			NodeMiles:  nodeMiles,
			Mph:        mph,
		}
		res = append(res, car)
	}

	return res
}

/******************************
 *
 * Game Service GRPC API
 *
 ******************************/

func StartServer(ch chan int) {
	address := "0.0.0.0" + ":" + port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v...\n", address)

	var s *grpc.Server
	s = grpc.NewServer()

	g.RegisterRoadTripGameServer(s, &gameServer{})

	servErr := s.Serve(lis)
	if err != nil {
		panic(servErr)
	}

	ch <- 0
}

// UpsertCar creates or updates a car. Car should already have been created in the db.
func (*gameServer) UpsertCar(ctx context.Context, request *g.UpsertCarRequest) (*g.Empty, error) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	data.cars[request.Car.Id] = Car{
		Id:         request.Car.Id,
		RouteId:    request.Car.Location.RouteId,
		RouteIndex: request.Car.Location.Index,
		NodeMiles:  float64(request.Car.Location.Miles),
		Mph:        float64(request.Car.Mph),
	}

	return &g.Empty{}, nil
}

// GetCarLocation returns the car location. Error if car not found.
func (*gameServer) GetCarLocation(ctx context.Context, request *g.GetCarLocationRequest) (*g.Location, error) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	c, ok := data.cars[request.CarId]
	if !ok {
		return &g.Location{}, errors.New("Car not found")
	}

	l := g.Location{
		RouteId: c.RouteId,
		Index:   int32(c.RouteIndex),
		Miles:   float32(c.NodeMiles),
	}

	return &l, nil
}

/******************************
 *
 * Update Loop
 *
 ******************************/

// mainLoop updates the world every second
func mainLoop(ch chan int) {
	const updateInterval = time.Second
	const saveInterval = 10 * time.Second

	u := time.NewTicker(updateInterval)
	s := time.NewTicker(saveInterval)

	for {
		select {
		case <-u.C:
			update()
		case <-s.C:
			s.Stop()
			<-save()
			s.Reset(saveInterval)
		}
	}
}

// update makes all changes to world state
// Currently that means moving cars around.
func update() {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	fmt.Printf("Updating...\n")
	now := time.Now()
	diff := now.Sub(lastTick)
	lastTick = time.Now()
	for i, _ := range data.cars {
		car := data.cars[i]
		if car.Mph == 0 {
			continue
		}

		route := data.routes[car.RouteId]
		node := route.Nodes[car.RouteIndex]
		car.NodeMiles += car.Mph * (diff.Seconds() / 3600.0)
		excessMiles := car.NodeMiles - node.Miles
		for excessMiles > 0 {
			// move to next node
			car.RouteIndex++
			node = route.Nodes[car.RouteIndex]
			car.NodeMiles = excessMiles
			if car.RouteIndex == int32(len(route.Nodes)-1) {
				// end of the line
				car.Mph = 0
				break
			}
			excessMiles = car.NodeMiles - node.Miles
		}

		data.cars[i] = car

		// debug
		fmt.Printf("%+v\n", car)
	}
}

// save persists in-memory data to the database for safe keeping
func save() <-chan int32 {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	fmt.Printf("Saving...\n")
	c := make(chan int32)
	go func() {
		defer close(c)
		ctx := context.Background()
		for _, car := range data.cars {
			exec, err := client.Car.FindUnique(
				db.Car.ID.Equals(car.Id),
			).Update(
				db.Car.NodeMiles.Set(float64(car.NodeMiles)),
				db.Car.RouteIndex.Set(int(car.RouteIndex)),
			).Exec(ctx)
			if err != nil {
				return
			}
			fmt.Printf("%+v", exec)
		}
		c <- 1
	}()
	return c
}
