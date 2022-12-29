package main

import (
	"context"
	"log"
	"net"

	connection "mock-grpc/connection"
	"mock-grpc/pkg/database"
	server "mock-grpc/pkg/server"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func myInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	resp, err := handler(ctx, req)
	return resp, err
}

// declaring the port number
const (
	port = ":50051"
)

func main() {
	//creating gorm instance
	connection.DBinit()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//connecting gorm with grpc
	connection, err := gorm.Open("postgres", "user=postgres password=root dbname=zomato sslmode=disable")
	utils.CheckError(err)
	defer connection.Close()

	//creating a new server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(myInterceptor))
	zomatoServer := &server.ZomatoServer{
		Db: database.DBClient{
			Db: connection,
		},
	}
	pb.RegisterZomatoDatabaseCrudServer(grpcServer, zomatoServer)
	//if connection fails
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
