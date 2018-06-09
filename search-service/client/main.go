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

	pb "github.com/timothympace/tradesy-go-hackathon/search-service/search"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9092"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearchApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	filter := &pb.SearchFilter{Name: "Red"}
	items, err := c.GetItemByName(ctx, filter)
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
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
