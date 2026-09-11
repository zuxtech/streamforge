/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package http

import (
	"net/http"
	"time"
)

type LivenessResponse struct {
	Status string `json:"status"`
	Uptime uint64 `json:"uptime_seconds"`
}

type ReadinessResponse struct {
	Status string `json:"status"`
}

type Server struct {
	startTime time.Time
}

func NewServer() *Server {
	return &Server{
		startTime: time.Now(),
	}
}

func (s *Server) livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")

	writeJSON(w, http.StatusOK, LivenessResponse{
		Status: "ok",
		Uptime: uint64(time.Since(s.startTime).Seconds()),
	})
}

func (s *Server) readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Dependency checks will go here.
	// PostgreSQL, Redis, etc.

	w.Header().Set("Cache-Control", "no-cache")

	writeJSON(w, http.StatusOK, ReadinessResponse{
		Status: "ok",
	})
}