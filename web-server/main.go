package main

import (
	"encoding/json"
	"net/http"
	"log"

	itemPb "../item-service/item"
	userPb "../user-service/user"
	searchPb "../search-service/search"
	"time"
	"context"
	"google.golang.org/grpc"
)

const (
	addressItem   = "localhost:9090"
	addressUser   = "localhost:9091"
	addressSearch = "localhost:9092"
)

var itemApiClient itemPb.ItemApiClient
var userApiClient userPb.UserApiClient
var searchApiClient searchPb.SearchApiClient

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

func deleteItem(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := &itemPb.Id{}
	err := json.NewDecoder(req.Body).Decode(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	itemApiClient.DeleteItem(ctx, id)
	w.WriteHeader(http.StatusOK)
}

func allUsers(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var users []*userPb.User
	userStream, err := userApiClient.GetUser(ctx, &userPb.UserFilter{Id: ""})
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
	}

	for {
		user, err := userStream.Recv()
		if err != nil {
			break
		}
		users = append(users, user)
	}


	itemJSON, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Write(itemJSON)
}

func addUser(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user := &userPb.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}


	userApiClient.AddUser(ctx, user)
	w.WriteHeader(http.StatusCreated)
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := &userPb.UserFilter{}
	err := json.NewDecoder(req.Body).Decode(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userApiClient.DeleteUser(ctx, id)
	w.WriteHeader(http.StatusOK)
}

func getItemByName(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	filter := &searchPb.SearchFilter{}
	err := json.NewDecoder(req.Body).Decode(filter)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	itemStream, err := searchApiClient.GetItemByName(ctx, filter)


	var items []*searchPb.Item
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

func handlePublic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../" + r.URL.Path[1:])
}

func main() {
	conn, err := grpc.Dial(addressItem, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	itemApiClient = itemPb.NewItemApiClient(conn)

	conn, err = grpc.Dial(addressUser, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userApiClient = userPb.NewUserApiClient(conn)

	conn, err = grpc.Dial(addressSearch, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	searchApiClient = searchPb.NewSearchApiClient(conn)

	http.HandleFunc("/allItems", allItems)
	http.HandleFunc("/allUsers", allUsers)
	http.HandleFunc("/addItem", addItem)
	http.HandleFunc("/addUser", addUser)
	http.HandleFunc("/deleteItem", deleteItem)
	http.HandleFunc("/deleteUser", deleteUser)
	http.HandleFunc("/getItemByName", getItemByName)
	http.HandleFunc("/static/", handlePublic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
