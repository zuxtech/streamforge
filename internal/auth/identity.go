/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package auth

// Identity represents an authenticated StreamForge identity.

import "context"

type contextKey string

const identityContextKey contextKey = "auth_identity"

func IdentityFromContext(ctx context.Context) (*Identity, bool) {
	identity, ok := ctx.Value(identityContextKey).(*Identity)
	return identity, ok
}

func withIdentity(
	ctx context.Context,
	identity *Identity,
) context.Context {
	return context.WithValue(
		ctx,
		identityContextKey,
		identity,
	)
}
