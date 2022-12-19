package main

import (
	"zomato/connection"

	_ "github.com/lib/pq"
)

func main() {

	connection.DBinit()

}
