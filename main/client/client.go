package main

import (
	"context"
	"log"
	"time"

	"mock-grpc/pkg/client/agent"
	pb "mock-grpc/zomato-proto"

	"google.golang.org/grpc"
)

// declaring port for the client to run on
const (
	address = "localhost:50051"
)

func main() {

	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection Failed", err.Error())
	}
	defer connection.Close()

	C := pb.NewZomatoDatabaseCrudClient(connection)

	Ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// user.CreateUser(C, Ctx)
	// user.UpdateUser(C, Ctx, "Akshith", "bharadwaj@gmail.com")
	// user.GetUserOrders(C, Ctx)
	// user.GetUsers(C, Ctx)

	// restaurants.CreateRestaurant(C, Ctx)
	// restaurants.GetRestaurantMenu(C, Ctx, 1)
	// restaurants.UpdateRestaurant(C, Ctx, "mehfil", "banjara hills")

	// menu.CreateDish(C, Ctx)
	// menu.UpdateDish(C, Ctx, "lays", 500)
	// menu.DeleteDish(C, Ctx, "Tikka")

	// agent.CreateAgent(C, Ctx)
	// agent.UpdateAgentStatus(C, Ctx, "Navdeep", false)
	agent.GetDeliveryOrders(C, Ctx, 1)

}
