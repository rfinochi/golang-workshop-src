package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	ID     int
	Title  string
	IsDone bool
}

var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func CreateItem(newItem Item) int {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return 0
	}

	stmt := `INSERT INTO items (title, isDone)
    VALUES(?, ?)`

	result, err := db.Exec(stmt, newItem.Title, newItem.IsDone)

	if err != nil {
		errorLog.Fatal(err)
		return 0
	}

	id, err := result.LastInsertId()

	if err != nil {
		errorLog.Fatal(err)
		return 0
	}

	return int(id)
}

func UpdateItem(item Item) {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return
	}

	stmt := `UPDATE items 
	set title = ?,	isDone = ?
	where id = ?`

	_, err = db.Exec(stmt, item.Title, item.IsDone, item.ID)

	if err != nil {
		errorLog.Fatal(err)
	}
}

func GetItems() (items []Item) {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return nil
	}

	stmt := `SELECT id, title, isDone FROM items`

	rows, err := db.Query(stmt)

	if err != nil {
		errorLog.Fatal(err)
		return nil
	}

	defer rows.Close()

	items = []Item{}

	for rows.Next() {
		i := Item{}

		err := rows.Scan(&i.ID, &i.Title, &i.IsDone)
		if err != nil {
			errorLog.Fatal(err)
			return nil
		}

		items = append(items, i)
	}

	if err = rows.Err(); err != nil {
		errorLog.Fatal(err)
		return nil
	}

	return
}

func GetItem(id int) (item Item) {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return Item{}
	}

	stmt := `SELECT id, title, isDone FROM items
	where id = ?`

	i := Item{}

	err = db.QueryRow(stmt, id).Scan(&i.ID, &i.Title, &i.IsDone)

	if err != nil {
		errorLog.Fatal(err)
		return Item{}
	}
	return
}

func DeleteItem(id int) {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return
	}

	stmt := `DELETE FROM items where id = ?`

	_, err = db.Exec(stmt, id)

	if err != nil {
		errorLog.Fatal(err)
	}
}

func connnect() (*sql.DB, error) {
	//replace web with user
	//replace pass with password

	db, err := sql.Open("mysql", "web:pass@/toDo?parseTime=true")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func disconnect(db *sql.DB) {
	defer db.Close()
}
