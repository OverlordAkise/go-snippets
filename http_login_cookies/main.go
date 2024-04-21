package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var logins = sync.Map{}

func GenerateRandomKey() string {
	return strconv.Itoa(rand.Int())
}

func IsUserLoginCorrect(user, pw string) bool {
	if user == "" || pw == "" {
		return false
	}
	if user == "me" && pw == "me" {
		return true
	}
	return false
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		c, err := r.Cookie("key")
		if err != nil { //err is probably ErrNotFound
			fmt.Fprint(w, "You are not logged in!<br>")
			fmt.Fprint(w, "<a href='/login'>login</a>")
			return
		}
		name, ok := logins.Load(c.Value)
		if !ok {
			fmt.Fprint(w, "Your login cookie value is not existant! Please login again!<br>")
			fmt.Fprint(w, "<a href='/login'>login</a>")
			return
		}
		fmt.Fprintf(w, "Welcome, %s<br>", name)
		fmt.Fprint(w, "<a href='/logout'>logout</a>")
	})
	http.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<form action='/login' method='post'>")
		fmt.Fprint(w, "<input type='text' id='user' name='user'>")
		fmt.Fprint(w, "<input type='text' id='pass' name='pass'>")
		fmt.Fprint(w, "<input type='submit' value='Login'>")
		fmt.Fprint(w, "</form>")
	})
	http.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		pw := r.FormValue("pass")
		if _, err := r.Cookie("apikey"); err == nil {
			http.Redirect(w, r, "/", http.StatusFound) //already logged in, redirecting to home
			return
		}
		if !IsUserLoginCorrect(user, pw) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "wrong login information")
			return
		}
		key := GenerateRandomKey()
		logins.Store(key, user)
		ck := &http.Cookie{
			Name:     "key",
			Value:    key,
			MaxAge:   300,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, ck)

		http.Redirect(w, r, "/", http.StatusFound)
	})
	http.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("key")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ck := &http.Cookie{
			Name:   "key",
			Value:  "",
			MaxAge: -1,
		}
		logins.Delete(c.Value)
		http.SetCookie(w, ck)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	fmt.Println("Listening on :8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
