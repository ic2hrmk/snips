package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "net/http/pprof"
)

// Load HTTP benchmark
//go:generate go-wrk http://localhost:8080

func main() {
	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8080", httpServer))
}

func mainHandler(resp http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(struct{}{})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write(data)
}
