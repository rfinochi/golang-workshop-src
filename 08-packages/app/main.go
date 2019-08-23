package main

import (
	"fmt"
	"geometry"
	"github.com/rfinochi/golang-workshop-src/08-package/math"
)

func main() {
	var length, breadth = 10, 20
	fmt.Println("Area is", geometry.Area(length, breadth))
	fmt.Println("Perimeter is", geometry.Perimeter(length, breadth))
}
