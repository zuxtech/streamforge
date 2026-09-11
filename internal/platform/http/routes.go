/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Routes registers the StreamForge platform HTTP routes.

package http

import "net/http"

func (s *Server) registerRoutes(mux *http.ServeMux) {
	s.registerHealthRoutes(mux)
}