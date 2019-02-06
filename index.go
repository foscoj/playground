package main

import (
        "fmt"
        "hello"
        "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Test</h1>")
}
