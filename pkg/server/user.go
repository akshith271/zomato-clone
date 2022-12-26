package server

import (
	"context"
	"log"
	Mail "mock-grpc/mail"
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
	s.Db.CreateUser(newUser)
	Mail.Mail(newUser.Email)
	log.Printf("%v\n ", in.GetEmail())
	return &pb.User{
		Name:    in.GetName(),
		Phone:   in.GetPhone(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, nil
}

func (s *ZomatoServer) GetUsers(ctx context.Context, in *pb.VoidUserRequest) (*pb.AllUsers, error) {
	log.Printf("readUser method called from server side")
	totalUserData := []*pb.User{}
	totalUser, err := s.Db.GetUsers()
	for _, user := range totalUser {
		totalUserData = append(totalUserData, &pb.User{Name: user.Name, Address: user.Address, Email: user.Email, Phone: user.Phone})
	}
	return &pb.AllUsers{AllUsers: totalUserData}, err
}

func (s *ZomatoServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("updateUser method called from server side")
	currentUser := model.User{
		Name:    in.GetName(),
		Phone:   in.GetPhone(),
		Address: in.GetAddress(),
		Email:   in.GetEmail(),
	}
	err := s.Db.UpdateUser(currentUser)
	return &pb.User{
		Name:    in.GetName(),
		Phone:   in.GetPhone(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, err

}

func (s *ZomatoServer) GetUserOrders(ctx context.Context, in *pb.User) (*pb.UserOrder, error) {
	log.Printf("GetUserOrders method called from server side")
	result := []*pb.OrderItem{}
	user := model.User{}
	user.ID = uint(in.GetId())
	resultValue, err := s.Db.GetUserOrders(user)
	for _, value := range resultValue {
		result = append(result, &pb.OrderItem{
			DishId:   int64(value.DishId),
			Quantity: int64(value.Quantity),
		})
	}
	return &pb.UserOrder{OrderItems: result}, err
}
