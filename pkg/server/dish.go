package server

import (
	"context"
	"fmt"
	"log"
	model "mock-grpc/models"
	pb "mock-grpc/zomato-proto"
)

func (s *ZomatoServer) CreateDish(ctx context.Context, in *pb.NewDish) (*pb.Dish, error) {
	log.Printf("createDish method called from server side")
	newDish := model.Dish{
		Name:         in.GetName(),
		Description:  in.GetDescription(),
		Image:        in.GetImage(),
		Price:        int(in.GetPrice()),
		RestaurantId: uint(in.GetRestaurantId()),
		IsAvailable:  in.GetIsAvailable(),
		IsVeg:        in.GetIsVeg(),
		Cuisine:      in.GetCuisine(),
		Category:     in.GetCategory(),
	}
	err := s.Db.CreateDish(newDish)
	fmt.Println(newDish)
	return &pb.Dish{
		Name:         newDish.Name,
		Description:  newDish.Description,
		Image:        newDish.Image,
		Price:        int64(newDish.Price),
		RestaurantId: int64(newDish.RestaurantId),
		IsAvailable:  newDish.IsAvailable,
		IsVeg:        newDish.IsVeg,
		Cuisine:      newDish.Cuisine,
		Category:     newDish.Category}, err
}

func (s *ZomatoServer) UpdateDish(ctx context.Context, in *pb.Dish) (*pb.Dish, error) {
	log.Printf("updateDish method called from server side")
	updatedDish := model.Dish{
		Description: in.GetDescription(),
		Name:        in.GetName(),
		Price:       int(in.GetPrice()),
	}
	err := s.Db.UpdateDish(updatedDish)
	return &pb.Dish{
		Name:        updatedDish.Name,
		Price:       int64(updatedDish.Price),
		Description: updatedDish.Description,
	}, err
}

func (s *ZomatoServer) DeleteDish(ctx context.Context, in *pb.Dish) (*pb.VoidDishResponse, error) {
	log.Printf("Delete Dish method called from server side")
	err := s.Db.DeleteDish("Tikka")
	return &pb.VoidDishResponse{}, err
}
