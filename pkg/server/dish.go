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
	s.Db.Save(&newDish)
	fmt.Println(newDish)
	return &pb.Dish{
		Name:         in.GetName(),
		Description:  in.GetDescription(),
		Image:        in.GetImage(),
		Price:        in.GetPrice(),
		RestaurantId: in.GetRestaurantId(),
		IsAvailable:  in.GetIsAvailable(),
		IsVeg:        in.GetIsVeg(),
		Cuisine:      in.GetCuisine(),
		Category:     in.GetCategory()}, nil
}

func (s *ZomatoServer) UpdateDish(ctx context.Context, in *pb.Dish) (*pb.Dish, error) {
	log.Printf("updateDish method called from server side")
	s.Db.Model(&model.Dish{}).Where("name=?", in.GetName()).Update("price", in.GetPrice())
	return &pb.Dish{
		Name:        in.GetName(),
		Price:       in.GetPrice(),
		Description: in.GetDescription(),
	}, nil
}

func (s *ZomatoServer) DeleteDish(ctx context.Context, in *pb.Dish) (*pb.Dish, error) {
	log.Printf("Delete Dish method called from server side")
	s.Db.Where("name=?", in.GetName()).Delete(&model.Dish{})
	return &pb.Dish{
		Name:        in.GetName(),
		Price:       in.GetPrice(),
		Description: in.GetDescription(),
	}, nil
}
