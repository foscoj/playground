package main

import (
        "fmt"
        "net/http"
)

// Handler serves the Output to Output
func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Test 2</h1>")
}

func main(){
        // do nothing
}
