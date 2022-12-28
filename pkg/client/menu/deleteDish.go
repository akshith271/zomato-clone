package menu

import (
	"context"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func DeleteDish(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string) (*pb.VoidDishResponse, error) {
	message, err := c.DeleteDish(ctx, &pb.Dish{Name: newName})
	utils.CheckError(err)
	return message, nil
}
