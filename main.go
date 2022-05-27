package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    name, err := os.Hostname()
    if err != nil {
	panic(err)
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, from pod: %s\n", name)
    })

    http.ListenAndServe(":8080", nil)
}