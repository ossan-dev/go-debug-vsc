package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handleItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("managed:", r)
	time.Sleep(2 * time.Second) // emulates workload
	w.Write([]byte(fmt.Sprintln("handleItems invoked...")))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/items", handleItems)
	fmt.Println("ms-inventory starting on port", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
