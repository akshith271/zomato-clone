package menu

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func CreateDish(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	new_Dish, err := c.CreateDish(ctx, &pb.NewDish{
		Name:         "lays",
		Description:  "this is lays",
		Image:        "this is image",
		Price:        243,
		RestaurantId: 3,
		IsAvailable:  false,
		IsVeg:        true,
		Cuisine:      "american",
		Category:     "delight"})
	utils.CheckError(err)
	log.Printf("Dish Name: %v \n Email: %v", new_Dish.GetName(), new_Dish.GetPrice())
}
