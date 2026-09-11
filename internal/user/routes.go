/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Routes registers StreamForge user HTTP routes.

package user

import (
	"net/http"

	platformauth "github.com/zuxtech/streamforge/internal/platform/auth"
)

func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
	authenticator platformauth.Authenticator,
) {
	mux.Handle(
		"GET /v1/me",
		platformauth.Middleware(authenticator)(
			http.HandlerFunc(h.Me),
		),
	)
}

