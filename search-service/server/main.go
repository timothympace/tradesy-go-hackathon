/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pbItem "github.com/timothympace/tradesy-go-hackathon/item-service/item"
	pbSearch "github.com/timothympace/tradesy-go-hackathon/search-service/search"
	pbUser "github.com/timothympace/tradesy-go-hackathon/user-service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":9092"
	addressItem = "localhost:9090" //item
	addressUser = "localhost:9091" //user
)

type server struct{}

func (s *server) GetItemByName(filter *pbSearch.SearchFilter, stream pbSearch.SearchApi_GetItemByNameServer) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressItem, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Get instance of item api client
	itemApiClient := pbItem.NewItemApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// fucking magic here
	itemStream, err := itemApiClient.GetItems(ctx, &pbItem.Empty{})
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
	}

	//TODO Dont do this in prod, write a smarter getter
	// Set up a connection to the server.
	conn2, err2 := grpc.Dial(addressUser, grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	defer conn2.Close()

	// Get instance of item api client
	userApiClient := pbUser.NewUserApiClient(conn2)
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	for {
		item, err := itemStream.Recv()
		if err != nil {
			break
		}
		//fmt.Printf("Item: %v", item)

		if filter.Name != "" {
			if !strings.Contains(item.Name, filter.Name) {
				continue
			}
		}

		// calling the streaming API
		userFilter := &pbUser.UserFilter{Id: item.UserId}
		userStream, err3 := userApiClient.GetUser(ctx2, userFilter)
		if err3 != nil {
			log.Fatalf("Error on get user: %v", err3)
		}

		searchItem := &pbSearch.Item{}
		for {
			// Receiving the stream of data
			user, err3 := userStream.Recv()
			fmt.Println(user)
			if err3 == io.EOF {
				break
			}
			if item.UserId == "" { //getting late
				break
			}
			if err3 != nil {
				log.Fatalf("%v.GetUser(_) = _, %v", userApiClient, err3)
			}
			//log.Printf("User: %v", user)

			searchItem.UserId = user.Id
			searchItem.UserName = user.Name
		}
		//searchItem.UserName = "sadf"
		searchItem.UserId = item.UserId

		searchItem.Id = item.Id
		searchItem.Name = item.Name
		searchItem.Brand = item.Brand
		searchItem.Price = item.Price

		if err2 := stream.Send(searchItem); err2 != nil {
			return err2
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbSearch.RegisterSearchApiServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
