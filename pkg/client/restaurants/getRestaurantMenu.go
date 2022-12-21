package restaurants

import (
	"context"
	"fmt"
	"log"
	pb "mock-grpc/zomato-proto"
)

func GetRestaurantMenu(c pb.ZomatoDatabaseCrudClient, ctx context.Context, restaurantId int64) (*pb.Menu, error) {
	result, err := c.GetRestaurantMenu(ctx, &pb.RestaurantMenuRequest{
		RestaurantID: restaurantId,
	})
	if err != nil {
		log.Printf("error ochindi")
	}
	fmt.Println(result)
	return result, nil

}
