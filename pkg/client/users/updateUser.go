package users

import (
	"context"
	"mock-grpc/utils"
	pb "mock-grpc/zomato-proto"
)

func UpdateUser(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newEmail string) (*pb.User, error) {
	updated_user, err := c.UpdateUser(ctx, &pb.User{Name: newName, Email: newEmail})
	utils.CheckError(err)
	return updated_user, nil
}
