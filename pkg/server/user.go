package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	Mail "mock-grpc/mail"
	model "mock-grpc/models"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
	"net/mail"
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func (s *ZomatoServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("createUser method called from server side")
	// validate input
	if in.GetName() == "" || in.GetPhone() == "" || in.GetAddress() == "" || in.GetEmail() == "" {
		return nil, fmt.Errorf("invalid input")
	}
	if !isValidEmail(in.GetEmail()) {
		return nil, fmt.Errorf("invalid email")
	}

	newUser := model.User{
		Name:    in.GetName(),
		Phone:   in.GetPhone(),
		Address: in.GetAddress(),
		Email:   in.GetEmail(),
	}
	err := s.Db.CreateUser(newUser)
	utils.CheckError(err)
	log.Printf("%v\n ", in.GetEmail())
	Mail.Mail(in.GetEmail())
	return &pb.User{
		Id:      in.GetId(),
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

func (s *ZomatoServer) PlaceOrder(ctx context.Context, in *pb.OrderRequest) (*pb.Order, error) {
	log.Printf("PlaceOrder method called from server side")
	//get all agents who are active
	//pick one agent who is active at random
	//extract his id (agent_id)
	//get the id of the user who is calling this method (user_id)
	//create a new order with the agent_id and user_id and some restaurant_id
	totalAgentData := []*pb.Agent{}
	totalAgents, err := s.Db.GetAllAgents()
	// fmt.Print(len(totalAgents))
	for _, agent := range totalAgents {
		if agent.IsActive == true {
			totalAgentData = append(totalAgentData, &pb.Agent{
				Id:             int64(agent.ID),
				Name:           agent.Name,
				IsActive:       agent.IsActive,
				CurrentOrderId: int64(agent.CurrentOrderId),
			})
		}
	}
	utils.CheckError(err)
	index := rand.Intn(len(totalAgentData))
	randomAgentID := uint(totalAgentData[index].Id)
	userID := uint(in.GetUserID())
	restaurantID := uint(in.GetRestaurantID())
	newOrderRequest := model.Order{
		UserId:       userID,
		AgentId:      randomAgentID,
		RestaurantId: restaurantID,
	}
	errorResult := s.Db.CreateOrder(newOrderRequest)
	return &pb.Order{
		UserID:       in.GetUserID(),
		RestaurantID: in.GetRestaurantID(),
	}, errorResult
}
