package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const htmlBody = `<!DOCTYPE html>
<html>
<head><meta charset="utf-8"><title>Service Maintenance</title></head>
<body><h1>Under Maintenance</h1><p>This service is currently undergoing maintenance. Please try again later.</p></body>
</html>`

func main() {
	port := flag.Int("port", 8080, "listening port")
	retryAfter := flag.Int("retry-after", 3600, "Retry-After header value in seconds")
	flag.Parse()

	retryAfterStr := strconv.Itoa(*retryAfter)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Retry-After", retryAfterStr)
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, htmlBody)
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("servaway listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
