package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok9 github again")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/test", testHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
