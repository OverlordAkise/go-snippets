package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseGlob("main_files/*"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index", map[string]interface{}{
			"Title": "CrazyTitle",
			"Body":  "Hello, World!",
		})
		if err != nil {
			fmt.Println("ERROR executing template:", err)
		}
	})
	fmt.Println("Listening on :8091")
	fmt.Println(http.ListenAndServe(":8091", nil))
}
