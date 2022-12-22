package server

import (
	"context"
	"fmt"
	"log"
	model "mock-grpc/models"
	pb "mock-grpc/zomato-proto"
)

func (s *ZomatoServer) CreateAgent(ctx context.Context, in *pb.NewAgent) (*pb.Agent, error) {
	log.Printf("createAgent method called from server side")
	newAgent := model.Agent{
		Name:           in.GetName(),
		Phone:          in.GetPhone(),
		Address:        in.GetAddress(),
		Email:          in.GetEmail(),
		IsActive:       in.GetIsActive(),
		CurrentOrderId: int(in.GetCurrentOrderId()),
	}
	s.Db.Save(&newAgent)
	log.Printf("%v\n ", in.GetName())
	return &pb.Agent{
		Name:     in.GetName(),
		Phone:    in.GetPhone(),
		Address:  in.GetAddress(),
		Email:    in.GetEmail(),
		IsActive: in.GetIsActive(),
	}, nil
}

func (s *ZomatoServer) UpdateAgentStatus(ctx context.Context, in *pb.AgentStatus) (*pb.AgentStatus, error) {
	log.Printf("updateAgent method called from server side")
	s.Db.Model(&model.Agent{}).Where("name=?", in.GetName()).Update("is_active", in.GetIsActive())
	return &pb.AgentStatus{
		Name:     in.GetName(),
		IsActive: in.GetIsActive()}, nil
}

func (s *ZomatoServer) GetDeliveryOrders(ctx context.Context, in *pb.Agent) (*pb.AgentOrders, error) {
	log.Printf("GetUserOrders method called from server side")
	result := []*pb.OrderID{}
	s.Db.Model(&model.Order{}).Where("agent_id=?", in.GetId()).Find(&result)
	fmt.Println(result)
	for _, value := range result {
		result = append(result, &pb.OrderID{
			Id: value.Id,
		})
	}
	return &pb.AgentOrders{OrderID: result}, nil
}
