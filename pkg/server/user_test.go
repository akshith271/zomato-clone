package server

import (
	"mock-grpc/mocks"
	model "mock-grpc/models"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"

	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockOrderRequest = pb.OrderRequest{
	UserID:       1,
	RestaurantID: 1,
}
var mockOrder = pb.Order{
	OrderID:      1,
	AgentID:      6,
	UserID:       5,
	RestaurantID: 1,
}
var mockUser = model.User{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var mockPlacedAgent = model.Agent{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var newUser = pb.NewUser{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var userToBeUpdated = model.User{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var mockUpdatedUser = pb.User{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var mockOrderItem = model.OrderItem{
	DishId:   1,
	Quantity: 1,
}
var emptyMockUser = model.User{}
var user = pb.User{
	Id:      1,
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var userOrder = &pb.UserOrder{
	OrderItems: []*pb.OrderItem{
		{DishId: 1,
			Quantity: 1,
		},
	},
}

func TestCreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testUser := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateUser(mockUser).Return(nil)
	got, err := testUser.CreateUser(ctx, &newUser)
	utils.CheckError(err)
	expected := &pb.User{
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

func TestGetUsers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testUser := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().GetUsers().Return([]model.User{mockUser}, nil)
	got, err := testUser.GetUsers(ctx, &pb.VoidUserRequest{})
	utils.CheckError(err)
	result := got.AllUsers
	expected := []*pb.User{}
	expected = append(expected, &pb.User{
		Address: "ullamco labore aliqua",
		Email:   "esse",
		Name:    "sunt ut enim incididunt",
		Phone:   "velit",
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			result, expected)
	}
}

func TestUpdateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testUser := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateUser(userToBeUpdated).Return(nil)
	got, err := testUser.UpdateUser(ctx, &mockUpdatedUser)

	utils.CheckError(err)
	expected := &pb.User{
		Address: "ullamco labore aliqua",
		Email:   "esse",
		Name:    "sunt ut enim incididunt",
		Phone:   "velit",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestGetUserOrders(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testUser := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	emptyMockUser.ID = 1
	mockDb.EXPECT().GetUserOrders(emptyMockUser).Return([]model.OrderItem{mockOrderItem}, nil)
	got, err := testUser.GetUserOrders(ctx, &user)
	utils.CheckError(err)
	expected := userOrder
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			got, expected)
	}
}
