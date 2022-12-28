package restaurants

import (
	"context"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func UpdateRestaurant(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newAddress string) (*pb.Restaurant, error) {
	updatedRestaurant, err := c.UpdateRestaurant(ctx, &pb.Restaurant{Name: newName, Address: newAddress})
	utils.CheckError(err)
	return updatedRestaurant, nil
}
