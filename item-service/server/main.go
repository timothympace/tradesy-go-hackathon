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

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/timothympace/tradesy-go-hackathon/item-service/item"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9090"
)

type server struct{}

var items []*pb.Item

func (s *server) GetItem(ctx context.Context, in *pb.Id) (*pb.Item, error) {

	for _, item := range items {
		if item.Id == in.Id {
			return item, nil
		}
	}

	return nil, fmt.Errorf("item doesn't exist")
}

func (s *server) GetItems(in *pb.Empty, stream pb.ItemApi_GetItemsServer) error {
	for _, item := range items {
		if err := stream.Send(item); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) AddItem(ctx context.Context, in *pb.Item) (*pb.Id, error) {
	for _, item := range items {
		if item.Id == in.Id {
			return nil, fmt.Errorf("item with id '%v' already exists", in.Id)
		}
	}

	items = append(items, in)
	return &pb.Id{Id: in.Id}, nil
}

func (s *server) DeleteItem(ctx context.Context, in *pb.Id) (*pb.Id, error) {
	for idx, item := range items {
		if item.Id == in.Id {
			items[len(items)-1], items[idx] = items[idx], items[len(items)-1]
			items = items[:len(items)-1]
			return in, nil
		}
	}

	return nil, fmt.Errorf("could not find item with id %v", in.Id)
}

func (s *server) UpdateItem(ctx context.Context, in *pb.Item) (*pb.Id, error) {
	for idx, item := range items {
		if item.Id == in.Id {
			items[idx] = in
			return &pb.Id{Id: in.Id}, nil
		}
	}

	return nil, fmt.Errorf("could not find item with id %v to update", in.Id)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	item := &pb.Item{Name: "Red shirt", Id: "1000"}
	items = append(items, item)
	item = &pb.Item{Name: "Green shirt", Id: "1001"}
	items = append(items, item)
	item = &pb.Item{Name: "Yellow shirt", Id: "1002"}
	items = append(items, item)
	item = &pb.Item{Name: "Black pants", Id: "1003"}
	items = append(items, item)
	item = &pb.Item{Name: "Blue pants", Id: "1004"}
	items = append(items, item)
	item = &pb.Item{Name: "White shoes", Id: "1005"}
	items = append(items, item)
	item = &pb.Item{Name: "Blue shoes", Id: "1006"}
	items = append(items, item)

	s := grpc.NewServer()
	pb.RegisterItemApiServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
