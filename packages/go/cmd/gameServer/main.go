package main

import (
	"context"
	"errors"
	"fmt"
	. "github.com/brickshot/roadtrip-v2/internal/gameServer/grpc"
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

/******************************
 *
 * Main
 *
 ******************************/

func init() {
	lastTick = time.Now()
}

func main() {
	fmt.Printf("GameServer started...\n")
	c := make(chan int)
	go mainLoop(c)
	go StartServer(c)
	<-c
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
	for _,c := range data.cars {
		c.location.miles += int32(float64(c.mph) * (diff.Seconds() / 3600.0))
	}
}

