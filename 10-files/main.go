package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFromFile()
	writeToFile()
}

func readFromFile() {
	data, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func writeToFile() {
	data := []byte("This is some information!")
	err := ioutil.WriteFile("write_data.txt", data, 0666)
	if err != nil {
		fmt.Println("There has been an error:", err)
		return
	}
}