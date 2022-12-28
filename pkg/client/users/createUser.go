package users

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func CreateUser(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	new_user, err := c.CreateUser(ctx, &pb.NewUser{
		Name:    "Surya",
		Phone:   "9876543",
		Address: "habsiguda",
		Email:   "123@123.in"})
	utils.CheckError(err)
	log.Printf("User Name: %v \n Email: %v", new_user.GetName(), new_user.GetEmail())
}
