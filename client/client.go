package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "zomato/zomato-proto"

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

	c := pb.NewZomatoDatabaseCrudClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//create user from client
	new_user, err := c.CreateUser(ctx, &pb.NewUser{Name: "Madhav",
		Address: "habsiguda", Email: "123@123.in"})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("User Name: %v \n", new_user.GetName())

	//total user from client
	total_user, err := c.ReadUser(ctx, &pb.VoidUserRequest{})
	if err != nil {
		log.Printf("error getting all users")
	}
	//printing user details from the total_user
	for _, employee := range total_user.Users {
		fmt.Println(employee.GetName(), employee.GetEmail())
	}
	//updating the manager of an emp
	updated_user, err := c.UpdateUser(ctx, &pb.User{Name: "SuryaMadhav", Email: "sample@gmail.com"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(updated_user)

}
