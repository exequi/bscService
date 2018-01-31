package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Http service start up.....")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
