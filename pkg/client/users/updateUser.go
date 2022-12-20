package users

import (
	"context"
	pb "mock-grpc/zomato-proto"
)

func UpdateUser(c pb.ZomatoDatabaseCrudClient, ctx context.Context, newName string, newEmail string) (*pb.User, error) {
	updated_user, err := c.UpdateUser(ctx, &pb.User{Name: newName, Email: newEmail})
	if err != nil {
		panic(err.Error())
	}
	return updated_user, nil
}
