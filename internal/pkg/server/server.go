package server

import "net/http"

type Server struct {
	handler Handler
	config Config
}

func NewServer(config Config, handler Handler) (*Server, error) {
	serv := &Server{
		config: config,
		handler: handler,
	}

	return serv, nil
}

func Start(config Config, handler Handler) error {
	serv, err := NewServer(config, handler)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", serv.handler)
}


