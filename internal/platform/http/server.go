/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Server configures and serves the StreamForge HTTP API.

package http

import (
	"net/http"
	"time"

)


type Server struct {
	mux       *http.ServeMux
	handler   http.Handler
	startTime time.Time
	poweredBy string
}

type ServerOption func(*Server)

func WithPoweredBy(value string) ServerOption {
	return func(s *Server) {
		s.poweredBy = value
	}
}

func NewServer(options ...ServerOption) *Server {
	s := &Server{
		mux:       http.NewServeMux(),
		startTime: time.Now(),
	}

	for _, option := range options {
		option(s)
	}

	s.registerRoutes(s.mux)

	var handler http.Handler = s.mux

	handler = loggingMiddleware(handler)
	handler = poweredByMiddleware(s.poweredBy)(handler)
	handler = requestIDMiddleware(handler)

	s.handler = handler

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func (s *Server) Mux() *http.ServeMux {
	return s.mux
}