package server

import (
	"net/http"
	"time"

	"github.com/aliamerj/aircup/api/database"
	"github.com/aliamerj/aircup/config"
)

type Server struct {
	addr   string
	DB     database.Service
	Config config.Config
}

func NewServer(cfg config.Config, dburl string) (*Server, *http.Server, error) {
	db, err := database.New(dburl)
	if err != nil {
		return nil, nil, err
	}

	srv := &Server{
		addr:   cfg.Addr,
		DB:     db,
		Config: cfg,
	}

	httpServer := &http.Server{
		Addr:         cfg.Addr,
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv, httpServer, nil
}

func (s *Server) Close() error {
	if s.DB != nil {
		return s.DB.Close()
	}
	return nil
}
