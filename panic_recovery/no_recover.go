package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("After panic")
	}()

	panic("PANICME!")
}
