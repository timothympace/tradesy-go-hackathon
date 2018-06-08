package main

import (
	"encoding/json"
	"net/http"
	"log"

	itemPb "../item-service/item"
	"google.golang.org/grpc"
	"time"
	"context"
)

const (
	address     = "localhost:9090"
)

var itemApiClient itemPb.ItemApiClient

type Item struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Brand string  `json:"brand"`
	Size  string  `json:"size"`
	Color string  `json:"color"`
	Price float32 `json:"price"`
}

func allItems(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var items []*Item
	itemStream, err := itemApiClient.GetItems(ctx, &itemPb.Empty{})
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
	}

	for {
		item, err := itemStream.Recv()
		if err != nil {
			break
		}
		newItem := &Item{}
		newItem.Id = item.Id
		newItem.Name = item.Name
		newItem.Price = item.Price
		newItem.Brand = item.Brand
		items = append(items, newItem)
	}


	itemJSON, _ := json.Marshal(items)
	w.WriteHeader(http.StatusOK)
	w.Write(itemJSON)
}
/*
func addShirt(w http.ResponseWriter, req *http.Request) {
	var price float64

	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	price, err = strconv.ParseFloat(req.Form.Get("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item := Item{Brand: req.Form.Get("brand"), Size: req.Form.Get("size"), Color: req.Form.Get("color"), Price: price}
	itemList = append(itemList, item)
	w.WriteHeader(http.StatusCreated)
}

func writeShirts() {
	itemJSON, _ := json.Marshal(itemList)
	ioutil.WriteFile("items.json", itemJSON, os.ModePerm)
}

func init() {
	file, err := ioutil.ReadFile("items.json")
	if err == nil {
		json.Unmarshal(file, &itemList)
	}
}

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		writeShirts()
		os.Exit(0)
	}()
}*/

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	itemApiClient = itemPb.NewItemApiClient(conn)

	http.HandleFunc("/allItems", allItems)
	//http.HandleFunc("/addShirt", addShirt)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
