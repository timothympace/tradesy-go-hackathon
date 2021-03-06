package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/timothympace/tradesy-go-hackathon/user-service/user"
)

const (
	port = ":9091"
)

// server is used to implement user.UserServer.
type server struct {
	savedUsers []*pb.User
}

// AddUser creates a new User
func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	s.savedUsers = append(s.savedUsers, in)
	return &pb.UserResponse{Id: in.Id, Success: true}, nil
}

// UpdateUser creates a new User
func (s *server) UpdateUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	for index, user := range s.savedUsers {
		if in.Id != "" {
			if strings.Contains(user.Id, in.Id) {
				s.savedUsers[index] = in
				return &pb.UserResponse{Id: in.Id, Success: true}, nil
			}
		}
	}
	return &pb.UserResponse{Id: in.Id, Success: false}, nil
}

// GetUser returns all users by given filter
func (s *server) GetUser(filter *pb.UserFilter, stream pb.UserApi_GetUserServer) error {
	for _, user := range s.savedUsers {
		fmt.Printf("user.Id: %v, filter.Id: %v", user.Id, filter.Id)
		if filter.Id != "" {
			if !strings.Contains(user.Id, filter.Id) {
				continue
			}
		}
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

// DeleteUser deletes user by given filter
func (s *server) DeleteUser(ctx context.Context, filter *pb.UserFilter) (*pb.UserResponse, error) {
	for i, user := range s.savedUsers {
		if filter.Id != "" {
			if strings.Contains(user.Id, filter.Id) {
				s.savedUsers = append(s.savedUsers[:i], s.savedUsers[i+1:]...)
				return &pb.UserResponse{Id: filter.Id, Success: true}, nil
			}
		}
	}
	return &pb.UserResponse{Id: filter.Id, Success: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := &server{}
	user := &pb.User{
		Id:    "99",
		Name:  "Dmitriy Labitov",
		Email: "dmitriy@labitov.com",
		Phone: "444-444-55555",
		Addresses: []*pb.User_Address{
			&pb.User_Address{
				Street:            "444 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94106",
				IsShippingAddress: true,
			},
		},
	}
	server.savedUsers = append(server.savedUsers, user)

	user = &pb.User{
		Id:    "100",
		Name:  "Timothy Pace",
		Email: "timp@tradesy.com",
		Phone: "111-222-3344",
		Addresses: []*pb.User_Address{
			&pb.User_Address{
				Street:            "444 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94106",
				IsShippingAddress: true,
			},
		},
	}
	server.savedUsers = append(server.savedUsers, user)

	// Creates a new gRPC server
	s := grpc.NewServer()

	pb.RegisterUserApiServer(s, server)
	s.Serve(lis)
}
