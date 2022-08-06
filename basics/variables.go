package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
  display()
}

// Default values of int is 0 and bool is false
var i, j int = 1, 2

func display() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
