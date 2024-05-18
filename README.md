# Golang Code Snippets

This is a collection of knowledge and code examples for using go.

More documentation about each example can be found in their respective README.md file.

To learn more golang stuff visit [https://shira.at/blog/golang_notes.html](https://shira.at/blog/golang_notes.html)

## Misc

You can use 

    gofmt -w main.go

to auto format your main.go file with correct spacings.  
This changes

```go
type test struct {
    val1 string `json:"val1"`
    valuethree int `xml:"v3" json:"val3"`
}
```

to

```go
type test struct {
	val1       string `json:"val1"`
	valuethree int    `xml:"v3" json:"val3"`
}
```
