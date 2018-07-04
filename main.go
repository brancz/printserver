package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for headerName, values := range r.Header {
			for _, value := range values {
				fmt.Fprintf(os.Stdout, "%s: %s\n", headerName, value)
			}
		}
		fmt.Fprint(os.Stdout, "\n\n")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		os.Stdout.Write(b)

		fmt.Println("---")
	}))
}
