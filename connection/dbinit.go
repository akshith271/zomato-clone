package connection

import (
	"zomato/models"

	"github.com/jinzhu/gorm"
)

func DBinit() {
	db, err := gorm.Open("postgres", "user=postgres dbname=grpc password=root  sslmode=disable")
	CheckError(err)
	defer db.Close()
	db.AutoMigrate(&models.User{},
		&models.Agent{},
		&models.Dish{},
		&models.Hotel{},
		&models.Invoice{},
		&models.Order{},
		&models.OrderItem{},
	)
	//todo: create some static data for those tables

}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
