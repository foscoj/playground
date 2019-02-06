package main

import (
        "fmt"
        "net/http"
)

// Handler serves the Output to Output
func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Test</h1>")
}

func main(){
        fmt.Println("Hello")
}
