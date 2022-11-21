# Golang Code Snippets

This is a collection of knowledge and code examples for using go.

More documentation about each example can be found in their respective README.md file.

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

type test struct {
	val1       string `json:"val1"`
	valuethree int    `xml:"v3" json:"val3"`
}
```
