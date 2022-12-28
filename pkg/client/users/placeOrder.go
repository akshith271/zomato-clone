package users

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func PlaceOrder(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	newOrderResponse, err := c.PlaceOrder(ctx, &pb.OrderRequest{
		UserID:       2,
		RestaurantID: 1,
	})
	utils.CheckError(err)
	log.Printf("UserID: %v \n AgentID: %v", newOrderResponse.GetUserID(), newOrderResponse.GetAgentID())
}
