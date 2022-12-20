package menu

import (
	"context"
	pb "mock-grpc/zomato-proto"
)

func UpdateDish(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newPrice int64) (*pb.Dish, error) {
	updated_Dish, err := c.UpdateDish(ctx, &pb.Dish{Name: newName, Price: newPrice})
	if err != nil {
		panic(err.Error())
	}
	return updated_Dish, nil
}
