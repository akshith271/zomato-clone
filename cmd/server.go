package main

import (
	"context"
	"log"
	"net"

	"mock-grpc/pkg/database"
	server "mock-grpc/pkg/server"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func gRPCInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	return resp, err
}

// declaring the port number
const (
	port = ":54321"
)

func main() {
	//connecting gorm with grpc
	connection := database.DBinit()
	listener, err := net.Listen("tcp", port)
	utils.CheckError(err)
	defer connection.Close()
  log.Printf("Using port no %v", listener.Addr())
	//creating a new server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(gRPCInterceptor))
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
