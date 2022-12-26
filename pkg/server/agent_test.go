package server

import (
	"mock-grpc/mocks"
	model "mock-grpc/models"
	"reflect"

	"context"
	"testing"

	pb "mock-grpc/zomato-proto"

	"github.com/golang/mock/gomock"
)

var mockAgent = model.Agent{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var NewAgent = pb.NewAgent{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var agentStatusToBeUpdated = model.Agent{
	Name:     "Navdeep",
	IsActive: false,
}
var mockAgentStatusToBeUpdated = pb.AgentStatus{
	Name:     "Navdeep",
	IsActive: false,
}

func TestCreateAgent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testAgent := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateAgent(mockAgent).Return(nil)
	got, err := testAgent.CreateAgent(ctx, &NewAgent)
	if err != nil {
		panic(err)
	}
	expected := &pb.Agent{
		Address: "ullamco labore aliqua",
		Email:   "esse",
		Name:    "sunt ut enim incididunt",
		Phone:   "velit",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateAgentStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testAgent := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateAgentStatus(agentStatusToBeUpdated).Return(nil)
	got, err := testAgent.UpdateAgentStatus(ctx, &mockAgentStatusToBeUpdated)

	CheckError(err)
	expected := &pb.AgentStatus{
		Name:     "Navdeep",
		IsActive: false,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}
