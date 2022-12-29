package server

import (
	"context"
	"log"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
	"time"
)

// function to create new token on server
func (s *ZomatoServer) CreateToken(ctx context.Context, in *pb.User) (*pb.Token, error) {
	log.Printf("creating new token called")
	user, err := NewUserAuth(in.Name, in.Email)
	utils.CheckError(err)
	jwtManager := JWTManager{
		SecretKey:     "secret",
		TokenDuration: 15 * time.Minute,
	}
	token, err := jwtManager.Generate(user)
	utils.CheckError(err)
	if err != nil {
		return nil, err
	}
	return &pb.Token{Token: token}, nil
}
