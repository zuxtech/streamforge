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

	"github.com/google/uuid"
)


func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")

		if requestID == "" {
			requestID = "req_" + uuid.NewString()
		}

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)
	})
}


func poweredByMiddleware(value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if value != "" {
				w.Header().Set("X-Powered-By", value)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		// Add structured logging later.
		// Example:
		// method, path, request ID, duration, status, etc.
		_ = duration
	})
}