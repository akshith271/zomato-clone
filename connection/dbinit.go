package connection

import (
	"fmt"
	"mock-grpc/utils"

	"github.com/jinzhu/gorm"
)

func DBinit() {
	db, err := gorm.Open("postgres", "user=postgres dbname=zomato password=root  sslmode=disable")
	utils.CheckError(err)
	defer db.Close()
	fmt.Print("database created")
}
