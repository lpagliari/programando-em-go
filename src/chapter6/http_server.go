package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "%s", now)
	})

	http.ListenAndServe(":8081", nil)
}
