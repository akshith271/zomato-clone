package server

import (
	pb "mock-grpc/zomato-proto"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// connecting to the unimplemented methods via a struct
type ZomatoServer struct {
	pb.UnimplementedZomatoDatabaseCrudServer
	Db *gorm.DB //linking to db via this struct
}
