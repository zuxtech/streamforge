/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package auth provides HTTP routes for StreamForge authentication flows.

package auth

import (
	"net/http"
)

// RegisterRoutes registers authentication routes.
func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc(
		"GET /v1/auth/registration/flow",
		h.CreateRegistrationFlow,
	)

	 mux.HandleFunc(
        "POST /v1/auth/registration",
        h.CompleteRegistration,
    )
}


