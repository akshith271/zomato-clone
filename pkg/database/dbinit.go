package database

import (
	"fmt"
	"mock-grpc/utils"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func DBinit() *gorm.DB {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_DBNAME")
	db, err := gorm.Open("postgres", "user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	utils.CheckError(err)
	// defer db.Close()
	// db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Agent{},
	// 	&models.Dish{},
	// 	&models.Restaurant{},
	// 	&models.Order{},
	// 	&models.OrderItem{},
	// )
	//order
	// db.Model(&models.Order{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// db.Model(&models.Order{}).AddForeignKey("agent_id", "agents(id)", "CASCADE", "CASCADE")
	// db.Model(&models.Order{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	//dish
	// db.Model(&models.Dish{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	//orderItems
	// db.Model(&models.OrderItem{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// db.Model(&models.OrderItem{}).AddForeignKey("dish_id", "dishes(id)", "CASCADE", "CASCADE")
	// db.Model(&models.OrderItem{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")

	// db.Save(&models.Restaurant{
	// 	Name:     "Bawarchi",
	// 	Address:  "rtc x roads",
	// 	Timings:  "8-10",
	// 	Email:    "bawarchi@hyd.in",
	// 	Password: "12324",
	// 	IsActive: true,
	// })
	// db.Save(&models.Dish{
	// 	Name:         "kebab",
	// 	Description:  "this is kebab",
	// 	Image:        "this is kebab image",
	// 	Price:        167,
	// 	RestaurantId: 1,
	// 	IsAvailable:  true,
	// 	IsVeg:        true,
	// 	Cuisine:      "mughlai",
	// 	Category:     "delicacy",
	// })

	// db.Save(&models.Dish{
	// 	Name:         "Tikka",
	// 	Description:  "This is tikka",
	// 	Image:        "this is image",
	// 	Price:        153,
	// 	RestaurantId: 1,
	// 	IsAvailable:  true,
	// 	IsVeg:        true,
	// 	Cuisine:      "indian",
	// 	Category:     "delight",
	// })
	// db.Save(&models.Agent{
	// 	Name:           "Hemanth",
	// 	Phone:          "12312043284",
	// 	Address:        "nagole",
	// 	Email:          "hemanth@siubisd.com",
	// 	IsActive:       true,
	// 	CurrentOrderId: 1,
	// })
	// db.Save(&models.User{
	// 	Name:    "Akshith",
	// 	Phone:   "23423423",
	// 	Address: "Uppal",
	// 	Email:   "akshithbharadwaj@gmail.com",
	// })
	// db.Save(&models.User{
	// 	Name:    "GST",
	// 	Phone:   "45623423",
	// 	Address: "Habsiguda",
	// 	Email:   "gst@gmail.com",
	// })
	// db.Save(&models.OrderItem{
	// 	UserId:   1,
	// 	DishId:   1,
	// 	Quantity: 1,
	// 	OrderId:  1,
	// })
	// db.Save(&models.OrderItem{
	// 	UserId:   1,
	// 	DishId:   2,
	// 	Quantity: 1,
	// 	OrderId:  1,
	// })
	// db.Save(&models.Order{
	// 	UserId:       1,
	// 	AgentId:      1,
	// 	RestaurantId: 1,
	// })
	fmt.Print("database created")
	return db
}
