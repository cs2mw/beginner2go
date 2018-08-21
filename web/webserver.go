package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Scheme:" + r.URL.Scheme)
	fmt.Println("Path:" + r.URL.Path)
	fmt.Println(r.Form["param"])
	for k, v := range r.Form {
		fmt.Println("-------------华丽的分隔符----------------")
		fmt.Println("Key:" + k)
		fmt.Println("Value:" + strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello,I'm a webserver!")
}

func login(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	fmt.Println("method:", m)
	if m == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	// http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
