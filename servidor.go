package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
)

func main() {
	http.HandleFunc("/tempo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:02"))
	})
	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port),nil)
}