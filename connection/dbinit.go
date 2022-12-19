package connection

import (
	"fmt"
	"zomato/models"

	"github.com/jinzhu/gorm"
)

func DBinit() {
	db, err := gorm.Open("postgres", "user=postgres dbname=zomato password=root  sslmode=disable")
	CheckError(err)
	defer db.Close()
	db.AutoMigrate(&models.User{},
		&models.Agent{},
		&models.Dish{},
		&models.Restaurant{},
		&models.Order{},
		&models.OrderItem{},
	)
	// db.Save(&models.Restaurant{
	// 	Name:     "Bawarchi",
	// 	Address:  "rtc x roads",
	// 	Timings:  "8-10",
	// 	Email:    "bawarchi@hyd.in",
	// 	Password: "12324",
	// 	IsActive: true,
	// 	Dishes: []models.Dish{
	// 		{Name: "kebab",
	// 			Description:  "this is kebab",
	// 			Image:        "this is kebab image",
	// 			Price:        167,
	// 			RestaurantId: 1,
	// 			IsAvailable:  true,
	// 			IsVeg:        true,
	// 			Cuisine:      "mughlai",
	// 			Category:     "delicacy"},
	// 	},
	// 	Orders: []models.Order{},
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
	// 	CurrentOrderId: "1",
	// })
	// db.Save(&models.User{
	// 	Name:    "Akshith",
	// 	Phone:   "23423423",
	// 	Address: "Uppal",
	// 	Email:   "akshithbharadwaj@gmail.com",
	// 	Orders: []models.Order{
	// 		{UserId: 1,
	// 			AgentId:      1,
	// 			RestaurantId: 1,
	// 			OrderItems: []models.OrderItem{
	// 				{DishId: 1,
	// 					Quantity: 1,
	// 					OrderId:  1},
	// 			},
	// 		},
	// 	},
	// })
	fmt.Print("database created")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
