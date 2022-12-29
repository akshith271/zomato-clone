package users

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func CreateUser(c pb.ZomatoDatabaseCrudClient, ctx context.Context) (*pb.User, error) {
	newUser, err := c.CreateUser(ctx, &pb.NewUser{
		Address: "quis id deserunt cupidatat ad",
		Email:   "1234@bc.in",
		Name:    "culpa nisi",
		Phone:   "consequat incididunt laboris aute aliqua",
	})
	utils.CheckError(err)
	log.Printf("User Name: %v \n Email: %v", newUser.GetName(), newUser.GetEmail())
	return newUser, err
}
