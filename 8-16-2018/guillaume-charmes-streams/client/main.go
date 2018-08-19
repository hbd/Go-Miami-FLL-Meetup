package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main2() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", err)
	}
	fmt.Printf("End.\n")
}

func main() {
	r, w := io.Pipe()
	go func() {
		c
		defer func() {
			if err := w.Close(); err != nil {
				log.Fatalf("Close w pipe: %s", err)
			}
		}()
		for i := 0; i < 3; i++ {
			fmt.Fprintf(w, "Hello!! %s\n", time.Now())
			//			time.Sleep(1 * time.Second)
		}
	}()
	req, err := http.NewRequest("POST", "http://localhost:8081/handler2", r)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", err)
	}
	println("Wait")
	fmt.Printf("End.\n")
}
