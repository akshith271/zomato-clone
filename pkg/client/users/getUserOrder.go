package users

import (
	"context"
	"fmt"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func GetUserOrders(c pb.ZomatoDatabaseCrudClient, ctx context.Context) (*pb.UserOrder, error) {
	result, err := c.GetUserOrders(ctx, &pb.User{
		Id: 1,
	})
	utils.CheckError(err)
	fmt.Println(result)
	return result, nil

}
