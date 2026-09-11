/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Middleware authenticates HTTP requests and stores the identity in the request context.

package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const identityContextKey contextKey = "auth.identity"

func Middleware(authenticator Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := bearerToken(r)
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			identity, err := authenticator.GetIdentity(r.Context(), token)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				identityContextKey,
				identity,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func bearerToken(r *http.Request) (string, bool) {
	value := r.Header.Get("Authorization")

	if value == "" {
		return "", false
	}

	parts := strings.Fields(value)

	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", false
	}

	if parts[1] == "" {
		return "", false
	}

	return parts[1], true
}

func IdentityFromContext(ctx context.Context) (*Identity, bool) {
	identity, ok := ctx.Value(identityContextKey).(*Identity)
	return identity, ok
}