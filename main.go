package main

import(
	"net/http"
	"io"
	"log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is home page")
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}