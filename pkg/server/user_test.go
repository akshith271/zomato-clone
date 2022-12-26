package server

import (
	"mock-grpc/mocks"
	model "mock-grpc/models"
	pb "mock-grpc/zomato-proto"

	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockUser = model.User{
	Address: "ullamco labore aliqua",
	Email:   "esse",
	Name:    "sunt ut enim incididunt",
	Phone:   "velit",
}

var NewUser = pb.NewUser{
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

var MockUpdatedUser = pb.User{
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
var User = pb.User{
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func TestCreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testUser := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateUser(mockUser).Return(nil)
	got, err := testUser.CreateUser(ctx, &NewUser)
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
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
	got, err := testUser.UpdateUser(ctx, &MockUpdatedUser)

	CheckError(err)
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
	got, err := testUser.GetUserOrders(ctx, &User)
	if err != nil {
		panic(err)
	}
	expected := userOrder
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			got, expected)
	}
}

// 1. database go lo interface lo rayali
// 2. replace the database call with the mock call written in the interface
// 3. create _test file
// 4. replace the content of the functions
