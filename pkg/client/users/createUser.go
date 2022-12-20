package users

import (
	"context"
	"log"
	pb "mock-grpc/zomato-proto"
)

func CreateUser(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	//create user from client
	new_user, err := c.CreateUser(ctx, &pb.NewUser{
		Name:    "Surya",
		Phone:   "9876543",
		Address: "habsiguda",
		Email:   "123@123.in"})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("User Name: %v \n Email: %v", new_user.GetName(), new_user.GetEmail())
}
