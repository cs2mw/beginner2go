package main

import (
	"fmt"
	"net/http"
)

func handler(rw http.ResponseWriter,req *http.Request) {
	fmt.Fprintf(rw,"Hello %s!I'm a webapp.",req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/cs2mw",handler)
	http.ListenAndServe(":8080",nil)
}

