package main

import (
	"encoding/json"
	"net/http"
	"log"

	itemPb "../item-service/item"
	"time"
	"context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9090"
)

var itemApiClient itemPb.ItemApiClient

func allItems(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var items []*itemPb.Item
	itemStream, err := itemApiClient.GetItems(ctx, &itemPb.Empty{})
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
	}

	for {
		item, err := itemStream.Recv()
		if err != nil {
			break
		}
		items = append(items, item)
	}


	itemJSON, _ := json.Marshal(items)
	w.WriteHeader(http.StatusOK)
	w.Write(itemJSON)
}

func addItem(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	item := &itemPb.Item{}
	err := json.NewDecoder(req.Body).Decode(item)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}


	itemApiClient.AddItem(ctx, item)
	w.WriteHeader(http.StatusCreated)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../" + r.URL.Path[1:])
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	itemApiClient = itemPb.NewItemApiClient(conn)

	http.HandleFunc("/allItems", allItems)
	http.HandleFunc("/addItem", addItem)
	http.HandleFunc("/static/", handleStatic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
