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
	port := map[bool]string{true: os.Getenv("PORT"), false: "8080"}[ os.Getenv("PORT") != ""]

	http.ListenAndServe(fmt.Sprintf(":%s", port),nil)
}