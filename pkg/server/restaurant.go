package server

import (
	"context"
	"fmt"
	"log"
	model "mock-grpc/models"
	pb "mock-grpc/zomato-proto"
)

func (s *ZomatoServer) CreateRestaurant(ctx context.Context, in *pb.NewRestaurant) (*pb.Restaurant, error) {
	log.Printf("createRestaurant method called from server side")
	newRestaurant := model.Restaurant{
		Name:     in.GetName(),
		Address:  in.GetAddress(),
		Timings:  in.GetTimings(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		IsActive: in.GetIsActive(),
	}
	s.Db.CreateRestaurant(newRestaurant)
	fmt.Println(newRestaurant)
	return &pb.Restaurant{
		Name:     in.GetName(),
		Address:  in.GetAddress(),
		Timings:  in.GetTimings(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		IsActive: in.GetIsActive()}, nil
}

func (s *ZomatoServer) GetRestaurantMenu(ctx context.Context, in *pb.RestaurantMenuRequest) (*pb.Menu, error) {
	log.Printf("GetRestaurantOrders method called from server side")
	result := []*pb.Dish{}
	restaurant := model.Restaurant{}
	restaurant.ID = uint(in.GetRestaurantID())
	resultValue, err := s.Db.GetRestaurantMenu(restaurant)
	for _, value := range resultValue {
		result = append(result, &pb.Dish{
			Name:        value.Name,
			Description: value.Description,
			Image:       value.Image,
			IsAvailable: value.IsAvailable,
			IsVeg:       value.IsVeg,
			Cuisine:     value.Cuisine,
			Category:    value.Category,
		})
	}
	return &pb.Menu{Dishes: result}, err
}

func (s *ZomatoServer) UpdateRestaurant(ctx context.Context, in *pb.Restaurant) (*pb.Restaurant, error) {
	log.Printf("updateRestaurant method called from server side")
	restaurant := model.Restaurant{
		Name:    in.GetName(),
		Address: in.GetAddress(),
		Email:   in.GetEmail(),
	}
	s.Db.UpdateRestaurant(restaurant)
	return &pb.Restaurant{
		Name:    in.GetName(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, nil
}
