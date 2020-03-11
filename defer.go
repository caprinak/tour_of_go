package main

import "fmt"

/* the defer statement (with or without panic and recover) provides an unusual and powerful mechanism for control flow. */
func main() {
	defer fmt.Println("world")
	printHeader()
	defer printFooter()
	fmt.Println("hello")
}
