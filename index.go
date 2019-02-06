package main

import (
        "fmt"
        "net/http"
)

// Handler serves the Output to Output
func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Test 2</h1>")
}

/* // main function already declared by zeit.co now
func main(){
        // do nothing
}
}