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

	// count := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// if count >= 5 {
		// 	w.WriteHeader(500)
		// 	fmt.Println("users visit count hits the limit.")
		// 	return
		// }
		fmt.Println(time.Now().Format("2006-01-02 15:04:05 Monday") + " " + hostname + " " + r.UserAgent())
		fmt.Fprintln(w, hostname)

		// count++
	})
	fmt.Println("servcie starting...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
