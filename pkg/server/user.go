package server

import (
	"context"
	"log"
	model "mock-grpc/models"
	pb "mock-grpc/zomato-proto"
)

func (s *ZomatoServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("createUser method called from server side")
	newUser := model.User{
		Name:    in.GetName(),
		Phone:   in.GetPhone(),
		Address: in.GetAddress(),
		Email:   in.GetEmail(),
	}
	s.Db.Save(&newUser)
	log.Printf("%v\n ", in.GetName())
	return &pb.User{
		Name:    in.GetName(),
		Phone:   in.GetPhone(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, nil
}

func (s *ZomatoServer) GetUsers(ctx context.Context, in *pb.VoidUserRequest) (*pb.AllUsers, error) {
	log.Printf("readUser method called from server side")
	totalUser := []model.User{}
	totalUserData := []*pb.User{}
	s.Db.Find(&totalUser)
	for _, user := range totalUser {
		totalUserData = append(totalUserData, &pb.User{Name: user.Name, Address: user.Address, Email: user.Email})
	}
	return &pb.AllUsers{AllUsers: totalUserData}, nil
}

func (s *ZomatoServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("updateUser method called from server side")
	s.Db.Model(&model.User{}).Where("name=?", in.GetName()).Update("email", in.GetEmail())
	return &pb.User{
		Name:    in.GetName(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, nil
}

func (s *ZomatoServer) GetUserOrders(ctx context.Context, in *pb.User) (*pb.UserOrder, error) {
	log.Printf("GetUserOrders method called from server side")
	result := []*pb.OrderItem{}
	s.Db.Model(&model.OrderItem{}).Where("user_id=?", in.GetId()).Find(&result)
	for _, value := range result {
		result = append(result, &pb.OrderItem{
			DishId:   value.DishId,
			Quantity: value.Quantity,
		})
	}
	return &pb.UserOrder{OrderItems: result}, nil
}
