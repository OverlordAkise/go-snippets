package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.New("index").Parse(webTemplate))

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

var webTemplate = `<!DOCTYPE HTML>
<html>
    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        {{.Body}}
    </body>
</html>
`
