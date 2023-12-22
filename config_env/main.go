package main

import (
	"fmt"
	"os"
)

func main() {

	value, ok := os.LookupEnv("MYAPP_PORT")
	if !ok {
		fmt.Println("MYAPP_PORT is empty")
	} else {
		fmt.Println("MYAPP_PORT is:", value)
	}
}
