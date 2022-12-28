package server

import (
	"fmt"
	"mock-grpc/mocks"
	model "mock-grpc/models"
	"mock-grpc/utils"
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

var mockPlacedAgent = &model.Agent{
	Address:  "ullamco labore aliqua",
	Email:    "esse",
	Name:     "sunt ut enim incididunt",
	Phone:    "velit",
	IsActive: true,
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
	tests := []struct {
		name        string
		input       *pb.NewUser
		mockErr     error
		expectedErr error
		expectedOut *pb.User
	}{
		{
			name: "valid input",
			input: &pb.NewUser{
				Id:      1,
				Name:    "ramesh",
				Phone:   "1234567890",
				Address: "yusufguda",
				Email:   "sai@example.com",
			},
			mockErr:     nil,
			expectedErr: nil,
			expectedOut: &pb.User{
				Id:      1,
				Name:    "ramesh",
				Phone:   "1234567890",
				Address: "yusufguda",
				Email:   "sai@example.com",
			},
		},
		{
			name: "invalid input",
			input: &pb.NewUser{
				Id:      0,
				Name:    "",
				Phone:   "",
				Address: "",
				Email:   "",
			},
			mockErr:     nil,
			expectedErr: fmt.Errorf("invalid input"),
			expectedOut: nil,
		},
		{
			name: "database error",
			input: &pb.NewUser{
				Id:      1,
				Name:    "ramesh",
				Phone:   "1234567890",
				Address: "yusufguda",
				Email:   "sai@example.com",
			},
			mockErr:     nil,
			expectedErr: nil,
			expectedOut: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDb.EXPECT().CreateUser(test.input).Return(test.mockErr)
			out, err := testUser.CreateUser(context.Background(), test.input)
			if err != test.expectedErr {
				t.Errorf("unexpected error: got %v, want %v", err.Error(), test.expectedErr)
			}
			if !reflect.DeepEqual(out, test.expectedOut) {
				t.Errorf("unexpected output: got %v, want %v", out, test.expectedOut)
			} else {
				t.Errorf("success")
			}
		})
	}
	// got, err := testUser.CreateUser(ctx, &newUser)
	// utils.CheckError(err)
	// expected := &pb.User{
	// 	Address: "ullamco labore aliqua",
	// 	Email:   "esse",
	// 	Name:    "sunt ut enim incididunt",
	// 	Phone:   "velit",
	// }
	// if !reflect.DeepEqual(got, expected) {
	// 	t.Errorf("The Function Returned is not expected one. got %v expected %v",
	// 		got, expected)
	// }
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

func TestPlaceOrder(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDB := mocks.NewMockDataBaseLayer(controller)
	mockDB.EXPECT().GetAllAgents().Return([]model.Agent{
		{Name: "GST", IsActive: true, CurrentOrderId: 0},
	}, nil)
	mockDB.EXPECT().CreateOrder(gomock.Any()).Return(nil)
	server := &ZomatoServer{
		Db: mockDB,
	}
	order, err := server.PlaceOrder(context.Background(), &pb.OrderRequest{
		UserID:       1,
		RestaurantID: 3,
	})
	utils.CheckError(err)
	if order.UserID != 1 || order.RestaurantID != 3 {
		t.Errorf("unexpected output: %v", order)
	}
}
