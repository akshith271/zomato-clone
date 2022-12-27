package database

import (
	"mock-grpc/models"

	"github.com/jinzhu/gorm"
)

type DataBaseLayer interface {
	//User interface calls
	CreateUser(models.User) error
	GetUsers() ([]models.User, error)
	UpdateUser(models.User) error
	GetUserOrders(models.User) ([]models.OrderItem, error)
	CreateOrder(models.Order) error
	//Agent interface calls
	CreateAgent(models.Agent) error
	UpdateAgentStatus(models.Agent) error
	GetAllAgents() ([]models.Agent, error)
	//Restaurant interface calls
	CreateRestaurant(models.Restaurant) error
	GetRestaurantMenu(models.Restaurant) ([]models.Dish, error)
	UpdateRestaurant(models.Restaurant) error
	//Dishes interface calls
	CreateDish(models.Dish) error
	UpdateDish(models.Dish) error
	DeleteDish(string) error
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) CreateUser(newUser models.User) error {
	db.Db.Save(&newUser)
	return nil
}

func (db DBClient) GetUsers() ([]models.User, error) {
	result := []models.User{}
	db.Db.Find(&result)
	return result, nil
}

func (db DBClient) UpdateUser(user models.User) error {
	db.Db.Model(&models.User{}).Where("name=?", user.Name).Update("email", user.Email)
	// db.Db.Updates(&user)
	return nil
}

func (db DBClient) GetUserOrders(user models.User) ([]models.OrderItem, error) {
	result := []models.OrderItem{}
	err := db.Db.Model([]models.OrderItem{}).Where("user_id=?", user.ID).Find(&result)
	return result, err.Error
}

func (db DBClient) CreateDish(newDish models.Dish) error {
	db.Db.Save(&newDish)
	return nil
}

func (db DBClient) UpdateDish(dish models.Dish) error {
	db.Db.Model(&models.Dish{}).Where("name=?", dish.Name).Update("price", dish.Price)
	return nil
}
func (db DBClient) DeleteDish(dishName string) error {
	DB := db.Db.Where("name=?", dishName).Delete(&models.Dish{})
	return DB.Error
}

func (db DBClient) CreateRestaurant(newRestaurant models.Restaurant) error {
	db.Db.Save(&newRestaurant)
	return nil
}
func (db DBClient) UpdateRestaurant(restaurant models.Restaurant) error {
	db.Db.Model(&models.Restaurant{}).Where("name=?", restaurant.Name).Update("address", restaurant.Address)
	return nil
}
func (db DBClient) GetRestaurantMenu(restaurant models.Restaurant) ([]models.Dish, error) {
	result := []models.Dish{}
	db.Db.Model(&models.Dish{}).Where("restaurant_id=?", restaurant.ID).Find(&result)
	return result, nil
}

func (db DBClient) CreateAgent(newAgent models.Agent) error {
	db.Db.Save(&newAgent)
	return nil
}

func (db DBClient) UpdateAgentStatus(agent models.Agent) error {
	db.Db.Model(&models.Agent{}).Where("name=?", agent.Name).Update("is_active", agent.IsActive)
	return nil
}

func (db DBClient) GetAllAgents() ([]models.Agent, error) {
	result := []models.Agent{}
	db.Db.Find(&result)
	return result, nil
}

func (db DBClient) CreateOrder(order models.Order) error {
	result := db.Db.Save(&order)
	return result.Error
}

// func (db DBClient) GetOrders() ([]models.Order, error) {
// 	result := []models.Order{}
// 	db.Db.Find(&result)
// 	return result, nil
// }
