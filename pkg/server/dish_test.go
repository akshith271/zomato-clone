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

var mockDish = model.Dish{
	Category:     "anim occaecat enim qui eu",
	Cuisine:      "Excepteur ut incididunt esse non",
	Description:  "exercitation do ea cillum sed",
	Image:        "sunt in enim commodo est",
	IsAvailable:  true,
	IsVeg:        true,
	Name:         "qui eiusmod sit",
	Price:        13078078312,
	RestaurantId: 8,
}

var NewDish = pb.NewDish{
	Category:     "anim occaecat enim qui eu",
	Cuisine:      "Excepteur ut incididunt esse non",
	Description:  "exercitation do ea cillum sed",
	Image:        "sunt in enim commodo est",
	IsAvailable:  true,
	IsVeg:        true,
	Name:         "qui eiusmod sit",
	Price:        13078078312,
	RestaurantId: 8,
}

var DishToBeUpdated = model.Dish{
	Description: "exercitation do ea cillum sed",
	Name:        "qui eiusmod sit",
	Price:       13078078312,
}

var MockUpdatedDish = pb.Dish{
	Description: "exercitation do ea cillum sed",
	Name:        "qui eiusmod sit",
	Price:       13078078312,
}

var deletedDish = pb.VoidDishResponse{}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func TestCreateDish(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testDish := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateDish(mockDish).Return(nil)
	got, err := testDish.CreateDish(ctx, &NewDish)
	if err != nil {
		panic(err)
	}
	expected := &pb.Dish{
		Category:     "anim occaecat enim qui eu",
		Cuisine:      "Excepteur ut incididunt esse non",
		Description:  "exercitation do ea cillum sed",
		Image:        "sunt in enim commodo est",
		IsAvailable:  true,
		IsVeg:        true,
		Name:         "qui eiusmod sit",
		Price:        13078078312,
		RestaurantId: 8,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Returned is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateDish(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testDish := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateDish(DishToBeUpdated).Return(nil)
	got, err := testDish.UpdateDish(ctx, &MockUpdatedDish)

	CheckError(err)
	expected := &pb.Dish{
		Description: "exercitation do ea cillum sed",
		Name:        "qui eiusmod sit",
		Price:       13078078312,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestDeleteDish(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBaseLayer(controller)
	testDish := ZomatoServer{Db: mockDb}
	ctx := context.Background()
	DishToBeDeleted := "Tikka"
	mockDb.EXPECT().DeleteDish(DishToBeDeleted).Return(nil)
	DishToBeDeleted = "Tikka"
	got, err := testDish.DeleteDish(ctx, &pb.Dish{})
	// model.CheckError(err)
	if err != nil {
		panic(err)
	}
	expected := &pb.VoidDishResponse{}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

// 1. database go lo interface lo rayali
// 2. replace the database call with the mock call written in the interface
// 3. create _test file
// 4. replace the content of the functions
