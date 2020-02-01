package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 202020
	fmt.Println(a)
}
