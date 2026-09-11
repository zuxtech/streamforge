package http

import "net/http"

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health/live", s.livenessHandler)
	mux.HandleFunc("GET /health/ready", s.readinessHandler)

	handler := http.Handler(mux)

	handler = loggingMiddleware(handler)
	handler = poweredByMiddleware(handler)
	handler = requestIDMiddleware(handler)

	return handler
}