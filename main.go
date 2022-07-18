package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")

}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")

}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
