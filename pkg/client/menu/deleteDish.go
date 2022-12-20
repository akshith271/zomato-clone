package menu

import (
	"context"
	pb "mock-grpc/zomato-proto"
)

func DeleteDish(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string) (*pb.VoidDishResponse, error) {
	message, err := c.DeleteDish(ctx, &pb.Dish{Name: newName})
	if err != nil {
		panic(err.Error())
	}
	return message, nil
}
