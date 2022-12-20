package main

import (
	"log"
	"net"

	connection "mock-grpc/connection"
	pb "mock-grpc/zomato-proto"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// declaring the port number
const (
	port = ":50051"
)

type ZomatoServer struct {
	pb.UnimplementedZomatoDatabaseCrudServer
	Db *gorm.DB //linking to db via this struct
}

func main() {

	//creating gorm instance
	connection.DBinit()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//connecting gorm with grpc
	connection, err := gorm.Open("postgres", "user=postgres password=root dbname=zomato sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	//creating a new server
	s := grpc.NewServer()

	pb.RegisterZomatoDatabaseCrudServer(s, ZomatoServer{
		Db: connection,
	})

	//if connection fails
	if err := s.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
