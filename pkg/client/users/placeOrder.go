package users

import (
	"context"
	"fmt"
	"log"
	pb "mock-grpc/zomato-proto"
)

func PlaceOrder(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	//create user from client
	fmt.Println("client called")
	newOrderResponse, err := c.PlaceOrder(ctx, &pb.OrderRequest{
		UserID:       1,
		RestaurantID: 1,
	})
	if err != nil {
		log.Fatal("place order function", err.Error())
	}
	log.Printf("UserID: %v \n AgentID: %v", newOrderResponse.GetUserID(), newOrderResponse.GetAgentID())
}
