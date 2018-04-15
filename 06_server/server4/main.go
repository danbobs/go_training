// Server3 returns a lissajous gif
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/danbobs/go_training/lissajous"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handles vy returning a lissajous gif with number of cycles determined by query string param
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if cycleStr, found := r.Form["cycles"]; found {
		cycles, err := strconv.Atoi(cycleStr[0])
		if err == nil {
			lissajous.CreateWithParams(w, cycles)
			return
		}
	}
	// create with default params
	lissajous.Create(w)
}
