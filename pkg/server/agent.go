package server

import (
	"context"
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
	err := s.Db.CreateAgent(newAgent)
	log.Printf("%v\n ", in.GetName())
	return &pb.Agent{
		Id:       int64(newAgent.ID),
		Name:     newAgent.Name,
		Phone:    newAgent.Phone,
		Address:  newAgent.Address,
		Email:    newAgent.Email,
		IsActive: newAgent.IsActive,
	}, err
}

func (s *ZomatoServer) UpdateAgentStatus(ctx context.Context, in *pb.AgentStatus) (*pb.AgentStatus, error) {
	log.Printf("updateAgent method called from server side")
	updatedAgent := model.Agent{
		Name:     in.GetName(),
		IsActive: in.IsActive,
	}
	err := s.Db.UpdateAgentStatus(updatedAgent)
	return &pb.AgentStatus{
		Name:     updatedAgent.Name,
		IsActive: updatedAgent.IsActive}, err
}
