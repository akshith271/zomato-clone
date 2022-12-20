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
	s.Db.Save(&newRestaurant)
	fmt.Println(newRestaurant)
	return &pb.Restaurant{
		Name:     in.GetName(),
		Address:  in.GetAddress(),
		Timings:  in.GetTimings(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		IsActive: in.GetIsActive()}, nil
}
