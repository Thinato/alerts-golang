package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port string) {
	http.HandleFunc("/health", healthCheck)
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":"+port, nil)
	fmt.Println("Running...")

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Health Check")
}
