// Fecth prints the content found at a url. Cheap "curl" knock-off
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, resp.Status)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s\n", body)
	}
}
