/*-- Create a new UTF-8 `toDo` database.
CREATE DATABASE toDo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using the `toDo` database.
USE toDo;

--Then copy and paste the following SQL statement to create a new snippets table to hold the text snippets for our application:

-- Create a `snippets` table.
CREATE TABLE items (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    isDone BOOLEAN
);

--Create mysql user
CREATE USER 'web'@'localhost';

GRANT SELECT, INSERT, UPDATE, DELETE ON toDo.* TO 'web'@'localhost';

-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

--

install mysql driver

go get github.com/go-sql-driver/mysql
*/

package main

import "fmt"

func main() {
	item := Item{
		Title:  "Test_1",
		IsDone: false,
	}

	itemID := CreateItem(item)
	fmt.Println(itemID)

	items := GetItems()
	fmt.Println(items)

	fmt.Println(GetItem(itemID))

	item.ID = itemID
	item.Title = "Test_1U"
	item.IsDone = true

	UpdateItem(item)
	fmt.Println(GetItem(itemID))

	DeleteItem(itemID)
}
