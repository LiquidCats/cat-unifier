package http

import (
	"cat-unifier/internal/kernel/common/library/repositiries"
	"fmt"
	"log"
	"net/http"
	"time"
)

type IServer interface {
	Serve(routes IHttpRoutes)
}

type server struct {
	cfg repositiries.IConfigRepository
}

func (s *server) Serve(routes IHttpRoutes) {
	port := s.cfg.Get("app.port", "8080")

	server := &http.Server{
		Handler:      routes.Init(),
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf(fmt.Sprintf("Starting on port [%s]", port))
	log.Fatal(server.ListenAndServe())
}

func NewServer(cfg repositiries.IConfigRepository) IServer {
	return &server{cfg: cfg}
}
