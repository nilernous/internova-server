package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from Go service")
    })

    addr := ":8080"
    fmt.Println("Go service listening on", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Println("server error:", err)
    }
}
