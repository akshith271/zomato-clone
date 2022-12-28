package menu

import (
	"context"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func UpdateDish(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newPrice int64) (*pb.Dish, error) {
	updated_Dish, err := c.UpdateDish(ctx, &pb.Dish{Name: newName, Price: newPrice})
	utils.CheckError(err)
	return updated_Dish, nil
}
