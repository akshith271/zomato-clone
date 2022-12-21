package users

import (
	"context"
	"fmt"
	"log"
	pb "mock-grpc/zomato-proto"
)

func GetUserOrders(c pb.ZomatoDatabaseCrudClient, ctx context.Context) (*pb.UserOrder, error) {
	result, err := c.GetUserOrders(ctx, &pb.User{
		Id: 1,
	})
	if err != nil {
		log.Printf("error ochindi")
	}
	fmt.Println(result)
	return result, nil

}
