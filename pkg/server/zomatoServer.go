package server

import (
	"mock-grpc/pkg/database"
	pb "mock-grpc/zomato-proto"

	_ "github.com/lib/pq"
)

// connecting to the unimplemented methods via a struct
type ZomatoServer struct {
	pb.UnimplementedZomatoDatabaseCrudServer
	Db database.DataBaseLayer //linking to db via this struct
}
