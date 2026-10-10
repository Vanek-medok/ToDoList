package server

import (
	"net/http"
	"strings"
)

func New(port string,
	addHandler http.Handler,
	helpHandler http.Handler,
	listHandler http.Handler,
	DeleteHandler http.Handler,
	DoneHandler http.Handler,
) *http.Server {
	mux := http.NewServeMux()

	mux.Handle("PATCH /tasks/{id}", DoneHandler)
	mux.Handle("DELETE /tasks/{id}", DeleteHandler)
	mux.Handle("GET /tasks", listHandler)
	mux.Handle("POST /tasks", addHandler)
	mux.Handle("GET /tasks/help", helpHandler)

	port = strings.TrimPrefix(port, ":")

	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
