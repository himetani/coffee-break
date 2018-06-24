package main

import (
	"flag"
	"fmt"
	"net/http"
)

//go:generate go-assets-builder -s="/static" -o assets.go static/

var isProd bool

func init() {
	flag.BoolVar(&isProd, "prod", false, "print version")
}

func main() {
	flag.Parse()

	if isProd {
		fmt.Println("serve from Assets")
		http.Handle("/", http.FileServer(Assets))
	} else {
		fmt.Println("serve from static dir")
		http.Handle("/", http.FileServer(http.Dir("./static")))
	}
	fmt.Println("Open :8080/")
	fmt.Println("Press ctrl+c to stop")
	http.ListenAndServe(":8080", nil)
}
