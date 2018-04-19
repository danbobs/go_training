// Server is a minimal "echo" web server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handles the request by echoing the path of the request back
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Path = %q\n", r.URL)
}
