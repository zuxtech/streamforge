/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Handler handles HTTP requests for StreamForge users.

package user

import (
	"net/http"

	platformauth "github.com/zuxtech/streamforge/internal/platform/auth"
	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
)

type Handler struct{}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	identity, ok := platformauth.IdentityFromContext(r.Context())

	if !ok {
		platformhttp.WriteJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
		return
	}

	platformhttp.WriteJSON(w, http.StatusOK, identity)
}
