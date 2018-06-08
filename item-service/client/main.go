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
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../item"
)

const (
	address     = "localhost:9090"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewItemApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.AddItem(ctx, &pb.Item{Name: "Tim", Id: "1"})
	if err != nil {
		log.Printf("could not add item: %v", err)
	}

	_, err = c.AddItem(ctx, &pb.Item{Name: "Dmitriy", Id: "2"})
	if err != nil {
		log.Printf("could not add item: %v", err)
	}

	_, err = c.AddItem(ctx, &pb.Item{Name: "Gustavo", Id: "3"})
	if err != nil {
		log.Printf("could not add item: %v", err)
	}

	_, err = c.DeleteItem(ctx, &pb.Id{Id: "2"})
	if err != nil {
		log.Printf("could not delete item: %v", err)
	}

	s, err2 := c.GetItem(ctx, &pb.Id{Id: "1"})
	if err2 != nil {
		log.Printf("could not get item: %v", err2)
	}
	log.Printf("Name: %s, Id: %s", s.Name, s.Id)

	items, err3 := c.GetItems(ctx, &pb.Empty{})
	if err3 != nil {
		log.Fatalf("could not get all items: %v", err3)
	}
	log.Println("All items:")
	for {
		item, err := items.Recv()
		if err != nil {
			break
		}
		log.Println(item)
	}

}
