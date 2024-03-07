package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handleInvoices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("managed:", r)
	time.Sleep(2 * time.Second) // emulates workload
	w.Write([]byte(fmt.Sprintln("handleInvoices invoked...")))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/invoices", handleInvoices)
	fmt.Println("ms-billing starting on port", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
