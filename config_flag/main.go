package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	//pointer,flagname,default,description
	flag.IntVar(&port, "port", 3000, "Port to listen on")

	var file string
	flag.StringVar(&file, "file", "", "File to read in")

	flag.Parse() //parse flags and set vars after defining them

	//If the -file is not used then it is empty, if mandatory you can do this:
	if file == "" {
		fmt.Println("ERROR: Please provide a file!")
		flag.PrintDefaults()
		return
	}

	fmt.Println("Port:", port)
	fmt.Println("File:", file)
}
