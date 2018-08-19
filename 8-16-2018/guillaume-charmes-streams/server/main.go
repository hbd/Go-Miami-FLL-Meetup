package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	for {
		fmt.Fprintf(w, "Hello world! %s\n", time.Now())
		flusher, ok := w.(http.Flusher)
		if !ok {
			log.Fatal("Not a flusher.")
		}
		flusher.Flush()
		time.Sleep(1 * time.Second)
	}
}

func handler2(w http.ResponseWriter, req *http.Request) {
	io.Copy(os.Stderr, req.Body)
	fmt.Printf("\n\nTHIS IS THE END!!!!")
}

func main() {
	http.HandleFunc("/handler1", handler)
	http.HandleFunc("/handler2", handler2)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
