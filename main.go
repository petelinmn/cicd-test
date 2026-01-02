package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok2")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
