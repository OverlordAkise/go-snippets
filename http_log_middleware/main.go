package main

import (
	"fmt"
	"net/http"
)

func rl(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("After Logging (recovered)")
				w.WriteHeader(500)
				fmt.Fprint(w, "ERR")
			} else {
				fmt.Println("After Logging")
			}
		}()
		fmt.Println("Before serving request")
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/", rl(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	}))
	http.HandleFunc("/panic", rl(func(w http.ResponseWriter, r *http.Request) {
		panic("panic with purpose")
	}))
	fmt.Println("Listening on :8090")
	fmt.Println(http.ListenAndServe(":8090", nil))
}
