package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Hotel struct {
	gorm.Model
	Name     string
	Address  string
	Timings  string
	Email    string
	Password string
	IsActive bool
	Dishes   *Dish
	Orders   *Order
}

type Dish struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	Price       int
	IsAvailable bool
	IsVeg       bool
	Cuisine     string
	Category    string
}

type User struct {
	gorm.Model
	Name    string
	Phone   string
	Address string
	Email   string
	Orders  *Order
}

type Agent struct {
	gorm.Model
	Name           string
	Phone          string
	Address        string
	Email          string
	IsActive       bool
	CurrentOrderId string
	Orders         *Order
}

type Order struct {
	gorm.Model
	UserId     uint
	AgentId    uint
	HotelId    uint
	InvoiceId  uint
	OrderItems *Order
}

type OrderItem struct {
	gorm.Model
	DishId   uint
	Quantity int
	OrderId  uint
}

type Invoice struct {
	gorm.Model
	UserId       uint
	PaymentMode  string
	Total        int
	TransactioId uint
}
