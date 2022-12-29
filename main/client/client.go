package main

import (
	"context"
	"fmt"
	"time"

	user "mock-grpc/pkg/client/users"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"

	"google.golang.org/grpc"
)

// declaring port for the client to run on
const (
	address = "localhost:50051"
)

func main() {

	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	utils.CheckError(err)
	defer connection.Close()

	C := pb.NewZomatoDatabaseCrudClient(connection)

	Ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newUser, err := user.CreateUser(C, Ctx)

	token, err := C.CreateToken(Ctx, newUser)
	fmt.Println(token)
	// token, err := server.CreateToken(C, Ctx, newUser)
	// fmt.Println(token)
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
	// user.PlaceOrder(C, Ctx)
	// agent.GetDeliveryOrders(C, Ctx, 1)

}
