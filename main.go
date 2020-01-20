package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	hostname, _ := os.Hostname()
	currentTime := time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(os.Stdout, currentTime.Format("2006-01-02 15:04:05 Monday")+hostname)
		fmt.Fprintln(w, currentTime.Format("2006-01-02 15:04:05 Monday")+hostname)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
