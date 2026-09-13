/*
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package http provides shared HTTP infrastructure for StreamForge.
package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type LivenessResponse struct {
	Status string `json:"status"`
	Uptime uint64 `json:"uptime_seconds"`
}

type ReadinessResponse struct {
	Status string `json:"status"`
}

func (s *Server) livenessHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Cache-Control", "no-cache")

	WriteJSON(w, http.StatusOK, LivenessResponse{
		Status: "ok",
		Uptime: uint64(time.Since(s.startTime).Seconds()),
	})
}

// readinessHandler reports whether StreamForge is ready to serve traffic.
func (s *Server) readinessHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// TODO: Check dependencies.
	//
	// Examples:
	// - PostgreSQL
	// - Redis
	// - Object storage
	// - Queue
	//
	// A failed dependency check should return a non-2xx status.

	w.Header().Set("Cache-Control", "no-cache")

	WriteJSON(w, http.StatusOK, ReadinessResponse{
		Status: "ready",
	})
}

// RegisterHealthRoutes registers health check routes on the given router.
func (s *Server) RegisterHealthRoutes(r chi.Router) {
	r.Get("/health/live", s.livenessHandler)
	r.Get("/health/ready", s.readinessHandler)
}
