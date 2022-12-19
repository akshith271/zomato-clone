package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name    string
	Phone   string
	Address string
	Email   string
	Orders  []Order
}
type Order struct {
	gorm.Model
	UserId       uint
	AgentId      uint
	RestaurantId uint
	OrderItems   []OrderItem
}

type OrderItem struct {
	gorm.Model
	DishId   int
	Quantity int
	OrderId  int
}
type Dish struct {
	gorm.Model
	Name         string
	Description  string
	Image        string
	Price        int
	RestaurantId uint
	IsAvailable  bool
	IsVeg        bool
	Cuisine      string
	Category     string
}
type Restaurant struct {
	gorm.Model
	Name     string
	Address  string
	Timings  string
	Email    string
	Password string
	IsActive bool
	Dishes   []Dish
	Orders   []Order
}

type Agent struct {
	gorm.Model
	Name           string
	Phone          string
	Address        string
	Email          string
	IsActive       bool
	CurrentOrderId string
}
