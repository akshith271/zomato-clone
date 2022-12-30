package server

import (
	"context"
	"log"
	"net"

	connection "zomato/connection"
	model "zomato/models"
	pb "zomato/zomato-proto"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// declaring the port number
const (
	port = ":54321"
)

// connecting to the unimplemented methods via a struct
type zomatoServer struct {
	pb.UnimplementedZomatoDatabaseCrudServer
	db *gorm.DB //linking to db via this struct
}

// server side implementation of creating user
func (s *zomatoServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("createUser method called from server side")
	newUser := model.User{
		Name:    in.GetName(),
		Address: in.GetAddress(),
		Email:   in.GetEmail(),
	}
	s.db.Save(&newUser)
	return &pb.User{
		Name:    in.GetName(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, nil
}

// server side implementation of reading user
func (s *zomatoServer) ReadUser(ctx context.Context, in *pb.VoidUserRequest) (*pb.AllUsers, error) {
	log.Printf("readUser method called from server side")
	totalUser := []model.User{}
	totalUserData := []*pb.User{}
	s.db.Find(&totalUser)
	for _, user := range totalUser {
		totalUserData = append(totalUserData, &pb.User{Name: user.Name, Address: user.Address, Email: user.Email})
	}
	return &pb.AllUsers{Users: totalUserData}, nil
}

// server side implementation of update user
// updates the manager of an user
func (s *zomatoServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("updateUser method called from server side")
	s.db.Model(&model.User{}).Where("name=?", in.GetName()).Update("email", in.GetEmail())
	return &pb.User{
		Name:    in.GetName(),
		Address: in.GetAddress(),
		Email:   in.GetEmail()}, nil
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

	pb.RegisterZomatoDatabaseCrudServer(s, &zomatoServer{
		db: connection,
	})

	//if connection fails
	if err := s.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
