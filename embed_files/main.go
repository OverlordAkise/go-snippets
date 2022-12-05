package main

import (
	"fmt"
	"net/http"
	"embed"
)

//go:embed testfile.txt
var tf string

//go:embed all:assets
var assets embed.FS

func main(){
	http.HandleFunc("/testfile",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "%q", tf)
	})
	http.Handle("/assets/", http.FileServer(http.FS(assets)))
    fmt.Println("starting")
	fmt.Println(http.ListenAndServe(":8080",nil))
    
    /*
    gofiber example:
    
    Rendering html templates via embed filesystem:
    
    --Remove the trailing folder, so that "assets/a.txt" becomes "a.txt"
    assetsFS, err := fs.Sub(assets, "assets")
    engine := html.NewFileSystem(http.FS(templateFS), ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    
    
    Hosting static files from embed filesystem:
    
    assetsFS, err := fs.Sub(assets, "assets")
    if err != nil {
        panic(err)
    }
    app.Use("/assets", filesystem.New(filesystem.Config{
        Root: http.FS(assetsFS),
    }))
    
    */
}
