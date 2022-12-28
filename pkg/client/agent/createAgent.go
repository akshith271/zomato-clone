package agent

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func CreateAgent(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	newAgent, err := c.CreateAgent(ctx, &pb.NewAgent{
		Name:           "Navdeep",
		Phone:          "765432",
		Address:        "ecil",
		Email:          "navdeep@gmail.com",
		IsActive:       false,
		CurrentOrderId: 1,
	})
	utils.CheckError(err)
	log.Printf("Agent Name: %v \n Email: %v", newAgent.GetName(), newAgent.GetEmail())
}
