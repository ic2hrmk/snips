package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TroubleShootServer struct{}

type CommonResponse struct {
	Timestamp time.Time `json:"timestamp"`
}

func NewTroubleShootServer() *TroubleShootServer {
	return &TroubleShootServer{}
}

func (rcv *TroubleShootServer) Run(port string) error {
	log.Printf("SERVER IS STRATING AT PORT [%s]...", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rcv.handleRoot)
	return http.ListenAndServe(":"+port, mux)
}

func (rcv *TroubleShootServer) handleRoot(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("-> SERVER ACCESSED FROM [%s] ON [%s], AGENT [%s]",
		req.RemoteAddr, req.RequestURI, req.UserAgent())

	successResponse := &CommonResponse{
		Timestamp: time.Now(),
	}

	data, err := json.MarshalIndent(successResponse, "  ", "  ")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Write(data)
}
