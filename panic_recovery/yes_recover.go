package main

import "fmt"

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic")
			fmt.Println(e)
		}
	}()

	panic("PANICME!")
}
