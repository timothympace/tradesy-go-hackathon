package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/timothympace/tradesy-go-hackathon/user-service/user"
)

const (
	address = "localhost:9091"
)

// AddUser calls the RPC method AddUser of UserApiClient
func addUser(client pb.UserApiClient, user *pb.User) {
	resp, err := client.AddUser(context.Background(), user)
	if err != nil {
		log.Fatalf("Could not create User: %v", err)
	}
	if resp.Success {
		log.Printf("A new User has been added with id: %s", resp.Id)
	}
}

// UpdateUser calls the RPC method UpdateUser of UserApiClient
func updateUser(client pb.UserApiClient, user *pb.User) {
	resp, err := client.UpdateUser(context.Background(), user)
	//log.Printf("%v, %v", resp, err)
	if err != nil {
		log.Fatalf("Could not update User: %v", err)
	}
	if resp.Success {
		log.Printf("A new User has been updated with id: %s", resp.Id)
	}
}

func deleteUser(client pb.UserApiClient, filter *pb.UserFilter) {
	resp, err := client.DeleteUser(context.Background(), filter)
	//log.Printf("resp: %v, err: %v", resp, err)
	if err != nil {
		log.Fatalf("Could not delete User: %v", err)
	}
	if resp.Success {
		log.Printf("A new has been deleted with id: %s", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getUser(client pb.UserApiClient, filter *pb.UserFilter) {
	// calling the streaming API
	stream, err :=
		client.GetUser(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get user: %v", err)
	}
	for {
		// Receiving the stream of data
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetUser(_) = _, %v", client, err)
		}
		log.Printf("User: %v", user)
	}
}
func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewUserApiClient(conn)

	// Create a new customer
	user := &pb.User{
		Id:    "101",
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.User_Address{
			&pb.User_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.User_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}
	addUser(client, user)

	// Create a new user
	user = &pb.User{
		Id:    "102",
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.User_Address{
			&pb.User_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}
	addUser(client, user)

	// Update a user
	user.Name = "Tim Rose"
	updateUser(client, user)

	// Create a new user
	user = &pb.User{
		Id:    "100",
		Name:  "Chuck Berry",
		Email: "c@b.com",
		Phone: "818-867-5309",
		Addresses: []*pb.User_Address{
			&pb.User_Address{
				Street:            "2nd Mission Street",
				City:              "Los Francisco",
				State:             "NV",
				Zip:               "88888",
				IsShippingAddress: false,
			},
		},
	}
	addUser(client, user)

	// Delete a user
	filter := &pb.UserFilter{Id: "101"}
	deleteUser(client, filter)

	// Filter with an empty Id
	filter = &pb.UserFilter{Id: ""}
	getUser(client, filter)
}
