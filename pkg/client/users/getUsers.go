package users

import (
	"context"
	"fmt"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func GetUsers(c pb.ZomatoDatabaseCrudClient, ctx context.Context) (*pb.AllUsers, error) {
	result, err := c.GetUsers(ctx, &pb.VoidUserRequest{})
	utils.CheckError(err)
	fmt.Println(result)
	return result, nil

}
