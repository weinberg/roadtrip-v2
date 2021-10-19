package main

import (
	"context"
	"errors"
	"fmt"
	. "github.com/brickshot/roadtrip-v2/internal/gameServer"
	. "github.com/brickshot/roadtrip-v2/internal/gameServer/grpc"
	db "github.com/brickshot/roadtrip-v2/internal/prisma"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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
	UnimplementedRoadTripGameServer
}

/******************************
 *
 * Data
 *
 ******************************/

// move this into a provider

type location struct {
	routeId string
	index   int32 // index of node in route if they were all laid out sequentially
	miles   int32
}

type car struct {
	id       string
	location location
	mph      int32
}

type gameData struct {
	cars map[string]car
}

var data = gameData{}
var lastTick time.Time
var client *db.PrismaClient
var routes []Route

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

	fmt.Printf("calling load routes")

	// Load routes
	<-loadRoutes()

	fmt.Printf("loaded routes")

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

func unmarshall(rs []db.RouteModel) []Route {
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

func loadRoutes() <-chan int32 {
	c := make(chan int32)
	go func() {
		defer close(c)
		ctx := context.Background()
		rs, _ := client.Route.FindMany().With(
			db.Route.Ways.Fetch().With(
				db.WaysOnRoutes.Way.Fetch().With(
					db.Way.Nodes.Fetch().With(
					  db.NodesOnWays.Node.Fetch(),
          ),
				),
			),
		).Exec(ctx)
		routes = unmarshall(rs)
		c <- 1
	}()
	return c
}

/******************************
 *
 * Game Service API methods
 *
 ******************************/

func StartServer(ch chan int) {
	address := "0.0.0.0" + ":" + port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v...\n", address)

	fmt.Printf("Connecting to data provider... unimplemented\n")

	var s *grpc.Server
	s = grpc.NewServer()

	RegisterRoadTripGameServer(s, &gameServer{})

	s.Serve(lis)
	ch <- 0
}

// UpsertCar creates or updates a car. Car should already have been created in the db.
func (*gameServer) UpsertCar(ctx context.Context, request *UpsertCarRequest) (*Empty, error) {
	data.cars[request.Car.Id] = car{
		id: request.Car.Id,
		location: location{
			routeId: request.Car.Location.RouteId,
			index:   request.Car.Location.Index,
			miles:   request.Car.Location.Miles,
		},
		mph: request.Car.Mph,
	}

	return &Empty{}, nil
}

//rpc GetCarLocation(GetCarLocationRequest) returns (Location) {
// GetCarLocation returns the car location. Error if car not found.
func (*gameServer) GetCarLocation(ctx context.Context, request *GetCarLocationRequest) (*Location, error) {
	c, ok := data.cars[request.CarId]
	if !ok {
		return &Location{}, errors.New("Car not found")
	}

	l := Location{
		RouteId: c.location.routeId,
		Index:   c.location.index,
		Miles:   c.location.miles,
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
	fmt.Println("Looping...")

	for {
		fmt.Println("Updating world...")
		update()
		time.Sleep(updateInterval)
	}
	ch <- 0
}

// update makes all changes to world state
// Currently that means moving cars around.
func update() {
	now := time.Now()
	diff := now.Sub(lastTick)
	for _, c := range data.cars {
		c.location.miles += int32(float64(c.mph) * (diff.Seconds() / 3600.0))
	}
}
