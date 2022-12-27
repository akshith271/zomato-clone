package server

import (
	"mock-grpc/mocks"
	"mock-grpc/models"
	model "mock-grpc/models"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"

	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockRestaurant = model.Restaurant{
	Address:  "in aliquip magna dolore ut",
	Email:    "labore in incididunt",
	IsActive: true,
	Name:     "Lorem irure dolor",
	Password: "Duis tempor et",
	Timings:  "pariatur commodo aliqua est mollit",
}

var emptyMockRestaurant = model.Restaurant{}

var newRestaurant = pb.NewRestaurant{
	Address:  "in aliquip magna dolore ut",
	Email:    "labore in incididunt",
	IsActive: true,
	Name:     "Lorem irure dolor",
	Password: "Duis tempor et",
	Timings:  "pariatur commodo aliqua est mollit",
}

var restaurantToBeUpdated = model.Restaurant{
	Address: "in aliquip magna dolore ut",
	Email:   "labore in incididunt",
	Name:    "Lorem irure dolor",
}

var mockUpdatedRestaurant = pb.Restaurant{
	Address: "in aliquip magna dolore ut",
	Email:   "labore in incididunt",
	Name:    "Lorem irure dolor",
}
var sampleMenu = models.Dish{
	Category:     "anim occaecat enim qui eu",
	Cuisine:      "Excepteur ut incididunt esse non",
	Description:  "exercitation do ea cillum sed",
	Image:        "sunt in enim commodo est",
	IsAvailable:  true,
	IsVeg:        true,
	Name:         "qui eiusmod sit",
	Price:        13078078312,
	RestaurantId: 1,
}

//	var mockDish = &pb.Dish{
//		Category:    "anim occaecat enim qui eu",
//		Cuisine:     "Excepteur ut incididunt esse non",
//		Description: "exercitation do ea cillum sed",
//		Image:       "sunt in enim commodo est",
//		IsAvailable: true,
//		IsVeg:       true,
//		Name:        "qui eiusmod sit",
//	}
var mockMenu = &pb.Menu{
	Dishes: []*pb.Dish{
		{
			Category:    "anim occaecat enim qui eu",
			Cuisine:     "Excepteur ut incididunt esse non",
			Description: "exercitation do ea cillum sed",
			Image:       "sunt in enim commodo est",
			IsAvailable: true,
			IsVeg:       true,
			Name:        "qui eiusmod sit",
		},
	},
}

var mockMenuRequest = pb.RestaurantMenuRequest{
	RestaurantID: 1,
}

func TestCreateRestaurant(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testRestaurant := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateRestaurant(mockRestaurant).Return(nil)
	got, err := testRestaurant.CreateRestaurant(ctx, &newRestaurant)
	utils.CheckError(err)
	expected := &pb.Restaurant{
		Address:  "in aliquip magna dolore ut",
		Email:    "labore in incididunt",
		IsActive: true,
		Name:     "Lorem irure dolor",
		Password: "Duis tempor et",
		Timings:  "pariatur commodo aliqua est mollit",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateRestaurant(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testRestaurant := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateRestaurant(restaurantToBeUpdated).Return(nil)
	got, err := testRestaurant.UpdateRestaurant(ctx, &mockUpdatedRestaurant)
	utils.CheckError(err)
	expected := &pb.Restaurant{
		Address: "in aliquip magna dolore ut",
		Email:   "labore in incididunt",
		Name:    "Lorem irure dolor",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestGetRestaurantMenu(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testRestaurant := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	emptyMockRestaurant.ID = 1
	mockDb.EXPECT().GetRestaurantMenu(emptyMockRestaurant).Return([]model.Dish{mockDish}, nil)
	got, err := testRestaurant.GetRestaurantMenu(ctx, &mockMenuRequest)
	utils.CheckError(err)
	expected := mockMenu
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			got, expected)
	}
}
