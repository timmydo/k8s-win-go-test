package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	endpoint = flag.String("endpoint", "https://www.bing.com", "the endpoint url to test")
	port     = flag.Int("port", 8080, "port to listen on")
)

func hello(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(*endpoint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		log.Printf("Error: %s\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Fprintf(w, "OK")
		log.Printf("Success calling: %s\n", *endpoint)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/healthz", hello)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
