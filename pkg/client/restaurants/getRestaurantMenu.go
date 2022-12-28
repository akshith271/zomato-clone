package restaurants

import (
	"context"
	"fmt"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func GetRestaurantMenu(c pb.ZomatoDatabaseCrudClient, ctx context.Context, restaurantId int64) (*pb.Menu, error) {
	result, err := c.GetRestaurantMenu(ctx, &pb.RestaurantMenuRequest{
		RestaurantID: restaurantId,
	})
	utils.CheckError(err)
	fmt.Println(result)
	return result, nil

}
