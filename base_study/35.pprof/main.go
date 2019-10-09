package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func sayHello(wr http.ResponseWriter, r *http.Request) {}

func main() {
	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
