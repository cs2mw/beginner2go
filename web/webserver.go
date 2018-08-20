package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Scheme:" + r.URL.Scheme)
	fmt.Println("Path:" + r.URL.Path)
	fmt.Println(r.Form["param"])
	for k, v := range r.Form {
		fmt.Println("Key:" + k)
		fmt.Println("Value:" + strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello,I'm a Go webserver!")
}

func main() {
	http.HandleFunc("/", sayHelloWorld)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
