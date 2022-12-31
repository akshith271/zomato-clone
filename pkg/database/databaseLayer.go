package database

import (
	"mock-grpc/models"

	"github.com/jinzhu/gorm"
)

type DataBaseLayer interface {
	//User
	CreateUser(models.User) error
	GetUsers() ([]models.User, error)
	UpdateUser(models.User) error
	GetUserOrders(models.User) ([]models.OrderItem, error)
	CreateOrder(models.Order) error
	//Agent
	CreateAgent(models.Agent) error
	UpdateAgentStatus(models.Agent) error
	GetAllActiveAgents() ([]models.Agent, error)
	//Restaurant
	CreateRestaurant(models.Restaurant) error
	GetRestaurantMenu(models.Restaurant) ([]models.Dish, error)
	UpdateRestaurant(models.Restaurant) error
	//Dishes
	CreateDish(models.Dish) error
	UpdateDish(models.Dish) error
	DeleteDish(string) error
}

type DBClient struct {
	Db *gorm.DB
}

// User calls
func (db DBClient) CreateUser(newUser models.User) error {
	err := db.Db.Save(&newUser)
	return err.Error
}
func (db DBClient) GetUsers() ([]models.User, error) {
	result := []models.User{}
	err := db.Db.Find(&result)
	return result, err.Error
}
func (db DBClient) UpdateUser(user models.User) error {
	err := db.Db.Model(&models.User{}).Where("name=?", user.Name).Update("email", user.Email)
	return err.Error
}

func (db DBClient) GetUserOrders(user models.User) ([]models.OrderItem, error) {
	result := []models.OrderItem{}
	err := db.Db.Preload("Orders.OrderItems").First(&user, "id = ?", user.ID)
	for _, order := range user.Orders {
		for _, orderItem := range order.OrderItems {
			result = append(result, orderItem)
		}
	}
	// err := db.Db.Model([]models.OrderItem{}).Where("user_id=?", user.ID).Find(&result)
	return result, err.Error
}

func (db DBClient) CreateDish(newDish models.Dish) error {
	err := db.Db.Save(&newDish)
	return err.Error
}

func (db DBClient) UpdateDish(dish models.Dish) error {
	err := db.Db.Model(&models.Dish{}).Where("name=?", dish.Name).Update("price", dish.Price)
	return err.Error
}
func (db DBClient) DeleteDish(dishName string) error {
	DB := db.Db.Where("name=?", dishName).Delete(&models.Dish{})
	return DB.Error
}

func (db DBClient) CreateRestaurant(newRestaurant models.Restaurant) error {
	err := db.Db.Save(&newRestaurant)
	return err.Error
}
func (db DBClient) UpdateRestaurant(restaurant models.Restaurant) error {
	err := db.Db.Model(&models.Restaurant{}).Where("name=?", restaurant.Name).Update("address", restaurant.Address)
	return err.Error
}
func (db DBClient) GetRestaurantMenu(restaurant models.Restaurant) ([]models.Dish, error) {
	result := []models.Dish{}
	err := db.Db.Model(&models.Dish{}).Where("restaurant_id=?", restaurant.ID).Find(&result)
	return result, err.Error
}

func (db DBClient) CreateAgent(newAgent models.Agent) error {
	err := db.Db.Save(&newAgent)
	return err.Error
}

func (db DBClient) UpdateAgentStatus(agent models.Agent) error {
	err := db.Db.Model(&models.Agent{}).Where("name=?", agent.Name).Update("is_active", agent.IsActive)
	return err.Error
}

func (db DBClient) GetAllActiveAgents() ([]models.Agent, error) {
	result := []models.Agent{}
	err := db.Db.Model(&models.Agent{}).Where("is_active=?", true).Find(&result)
	return result, err.Error
}

func (db DBClient) CreateOrder(order models.Order) error {
	err := db.Db.Save(&order)
	return err.Error
}
