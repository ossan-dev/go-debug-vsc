package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "unknown"
		}
		hostname, err := os.Hostname()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprintf("%q greets %q\n", hostname, name)))
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
