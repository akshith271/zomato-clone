package agent

import (
	"context"
	"fmt"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func UpdateAgentStatus(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newStatus bool) (*pb.AgentStatus, error) {
	updatedAgentStatus, err := c.UpdateAgentStatus(ctx, &pb.AgentStatus{Name: newName, IsActive: newStatus})
	utils.CheckError(err)
	fmt.Print(updatedAgentStatus)
	return updatedAgentStatus, nil
}
