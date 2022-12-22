package agent

import (
	"context"
	"log"
	pb "mock-grpc/zomato-proto"
)

func CreateAgent(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	//create Agent from client
	newAgent, err := c.CreateAgent(ctx, &pb.NewAgent{
		Name:           "Navdeep",
		Phone:          "765432",
		Address:        "ecil",
		Email:          "navdeep@gmail.com",
		IsActive:       false,
		CurrentOrderId: 1,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Agent Name: %v \n Email: %v", newAgent.GetName(), newAgent.GetEmail())
}
