/*
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

	"github.com/zuxtech/streamforge/internal/platform/id"
)

const (
	requestIDHeader = "X-Request-ID"
	poweredByHeader = "X-Powered-By"
)

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		requestID := r.Header.Get(requestIDHeader)

		if requestID == "" {
			requestID = id.New("req")
		}

		w.Header().Set(requestIDHeader, requestID)

		next.ServeHTTP(w, r)
	})
}

// poweredByMiddleware adds an X-Powered-By header to HTTP responses.
func poweredByMiddleware(value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			if value != "" {
				w.Header().Set(poweredByHeader, value)
			}

			next.ServeHTTP(w, r)
		})
	}
}

// loggingMiddleware records HTTP request timing.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		// TODO: Replace with structured logging.
		//
		// Fields should include:
		// - request ID
		// - method
		// - path
		// - status
		// - duration
		// - remote address
		//
		_ = duration
	})
}
