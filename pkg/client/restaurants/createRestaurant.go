package restaurants

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func CreateRestaurant(c pb.ZomatoDatabaseCrudClient, ctx context.Context) {
	new_Restaurant, err := c.CreateRestaurant(ctx, &pb.NewRestaurant{
		Name:     "mehfil",
		Address:  "jubilee hills",
		Timings:  "9-9",
		Email:    "mehfil@gmail.com",
		Password: "7654",
		IsActive: true})
	utils.CheckError(err)
	log.Printf("Restaurant Name: %v \n Email: %v", new_Restaurant.GetName(), new_Restaurant.GetEmail())
}
