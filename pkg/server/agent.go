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
	s.Db.CreateAgent(newAgent)
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
	agentToBeUpdated := model.Agent{
		Name:     in.GetName(),
		IsActive: in.IsActive,
	}
	s.Db.UpdateAgentStatus(agentToBeUpdated)
	return &pb.AgentStatus{
		Name:     in.GetName(),
		IsActive: in.GetIsActive()}, nil
}
