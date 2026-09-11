/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

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