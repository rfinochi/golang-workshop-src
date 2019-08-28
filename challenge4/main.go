package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type item struct {
	ID     int    `json:"Id"`
	Title  string `json:"Title"`
	IsDone bool   `json:"IsDone"`
}

type allItems []item

var items = allItems{
	{
		ID:     1,
		Title:  "Ir al workshop de Go",
		IsDone: true,
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the TodoAPI!")
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem item
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &newItem)
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newItem)
}

func getOneItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for _, singleItem := range items {
		if singleItem.ID == itemID {
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updatedItem item

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &updatedItem)

	for i, singleItem := range items {
		if singleItem.ID == itemID {
			singleItem.Title = updatedItem.Title
			singleItem.IsDone = updatedItem.IsDone
			items = append(items[:i], singleItem)
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for i, singleItem := range items {
		if singleItem.ID == itemID {
			items = append(items[:i], items[i+1:]...)
			fmt.Fprintf(w, "The item with ID %v has been deleted successfully", itemID)
			break
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home)
	router.HandleFunc("/todo", createItem).Methods("POST")
	router.HandleFunc("/todo", getAllItems).Methods("GET")
	router.HandleFunc("/todo/{id}", getOneItem).Methods("GET")
	router.HandleFunc("/todo/{id}", updateItem).Methods("PATCH")
	router.HandleFunc("/todo/{id}", deleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
