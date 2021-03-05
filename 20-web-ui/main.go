package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type toDoItem struct {
	ID      int
	Content string
	IsDone  bool
	Created time.Time
}

var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

var toDoItems = make([]toDoItem, 0, 0)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/home.page.tmpl",
		"./views/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		errorLog.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, toDoItems)
	if err != nil {
		errorLog.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		errorLog.Printf("error: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	todoItem := getToDoItem(id)

	if todoItem == nil {
		errorLog.Printf("Cant find ToDo Item with id %s", id)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	files := []string{
		"./views/show.page.tmpl",
		"./views/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		errorLog.Printf("error: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, todoItem)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func createForm(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/create.page.tmpl",
		"./views/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		errorLog.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	content := r.PostForm.Get("content")

	todoItem := toDoItem{
		ID:      rand.Intn(100),
		Content: content,
		IsDone:  false,
		Created: time.Now(),
	}

	toDoItems = append(toDoItems, todoItem)

	http.Redirect(w, r, fmt.Sprintf("/show/%d", todoItem.ID), http.StatusSeeOther)
}

func getToDoItem(id int) *toDoItem {
	for _, item := range toDoItems {
		if item.ID == id {
			return &item
		}
	}

	return nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/create", createForm).Methods("GET")
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/show/{id}", show).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	infoLog.Printf("Starting server on 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
