package users

import (
	"context"
	"fmt"
	"log"
	pb "mock-grpc/zomato-proto"
)

func GetUsers(c pb.ZomatoDatabaseCrudClient, ctx context.Context) (*pb.AllUsers, error) {
	result, err := c.GetUsers(ctx, &pb.VoidUserRequest{})
	if err != nil {
		log.Printf("error ochindi")
	}
	fmt.Println(result)
	return result, nil

}
