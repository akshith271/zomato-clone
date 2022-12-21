package restaurants

import (
	"context"
	pb "mock-grpc/zomato-proto"
)

func UpdateRestaurant(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newAddress string) (*pb.Restaurant, error) {
	updatedRestaurant, err := c.UpdateRestaurant(ctx, &pb.Restaurant{Name: newName, Address: newAddress})
	if err != nil {
		panic(err.Error())
	}
	return updatedRestaurant, nil
}
