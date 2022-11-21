package main

import (
	"fmt"
	"time"
)

func RecoverMe(id int, ch chan int) {
	if err := recover(); err != nil {
		fmt.Println("Recovered from panic:", err)
	}
	ch <- id
}

func Work(id int, ch chan int) {
	defer RecoverMe(id, ch)
	fmt.Println("Working on", id)
	time.Sleep(time.Duration(id) * time.Second)
	panic("PAUSE!")
}

func main() {
	ch := make(chan int)
	go Work(9, ch)
	go Work(3, ch)
	go Work(5, ch)
	go func() {
		for {
			i := <-ch
			fmt.Println("Worker died:", i)
			go Work(i, ch)
		}
	}()
	i := 0
	for {
		fmt.Println("Main", i)
		i += 10
		time.Sleep(10 * time.Second)
	}
}
