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
	"log"
	"net"
	"strings"
	"time"

	pbItem "github.com/timothympace/tradesy-go-hackathon/item-service/item"
	pbSearch "github.com/timothympace/tradesy-go-hackathon/search-service/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":9092"
	addressItem = "localhost:9090" //item
)

type server struct{}

var itemApiClient pbItem.ItemApiClient

func (s *server) GetItemByName(filter *pbSearch.SearchFilter, stream pbSearch.SearchApi_GetItemByNameServer) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressItem, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	itemApiClient := pbItem.NewItemApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	itemStream, err := itemApiClient.GetItems(ctx, &pbItem.Empty{})
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
	}

	for {
		item, err := itemStream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Item: %v", item)

		if filter.Name != "" {
			if !strings.Contains(item.Name, filter.Name) {
				continue
			}
		}

		searchItem := &pbSearch.Item{}
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
