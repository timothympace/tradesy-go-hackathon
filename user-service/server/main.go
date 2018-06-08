package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/timothympace/tradesy-go-hackathon/user-service/user"
)

const (
	port = ":50051"
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
		if filter.Id != "" {
			if !strings.Contains(user.Name, filter.Id) {
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
			if !strings.Contains(user.Name, filter.Id) {
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
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterUserApiServer(s, &server{})
	s.Serve(lis)
}
