package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type data struct {
	Num  int    `json:"Num"`
	Text string `json:"Text"`
}

func create(w http.ResponseWriter, r *http.Request) {
	var newData data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &newData)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newData)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The item with %v has been requested", mux.Vars(r)["num"])
}

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The all items has been requested")
}

func update(w http.ResponseWriter, r *http.Request) {
	var updatedData data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &updatedData)

	json.NewEncoder(w).Encode(updatedData)
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The item with %v has been deleted successfully", mux.Vars(r)["num"])
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				infoLog.Println(fmt.Errorf("%s", err))
				w.Header().Set("Connection", "close")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	logRequestMiddleware := alice.New(logRequest)
	recoverPanicMiddleware := alice.New(recoverPanic)

	router.Handle("/api", recoverPanicMiddleware.ThenFunc(create)).Methods("POST")
	router.HandleFunc("/api", getAll).Methods("GET")
	router.HandleFunc("/api/{num}", getOne).Methods("GET")
	router.Handle("/api/{num}", recoverPanicMiddleware.ThenFunc(update)).Methods("PATCH")
	router.Handle("/api/{num}", recoverPanicMiddleware.ThenFunc(delete)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", logRequestMiddleware.Then(router)))
}
