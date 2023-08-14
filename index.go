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

    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "api endpoint", name)
    })

    http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "payment page", name)
    })

    http.ListenAndServe(":8080", nil)
}
