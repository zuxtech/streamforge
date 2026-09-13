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
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers authentication routes.
func (h *Handler) RegisterRoutes(
	r chi.Router,
	identityProvider IdentityProvider,
) {
	r.Get(
		"/auth/registration/flow",
		h.CreateRegistrationFlow,
	)

	r.Post(
		"/auth/registration",
		h.CompleteRegistration,
	)

	r.With(
		Middleware(identityProvider),
	).Get(
		"/me",
		h.Me,
	)
}
