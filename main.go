package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(w, "<h1>HTTP Debug</h1>")
		fmt.Fprintln(w, "<pre>"+string(bytes)+"</pre>")

		log.Println("Serving request at /")
	})
	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
