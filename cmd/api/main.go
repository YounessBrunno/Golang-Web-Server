package main

import (
	"fmt"
	"log"
	"net/http"
)				

func main() {

  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
  })

}