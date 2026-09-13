/*
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package auth

import "context"

type Identity struct {
	ID    string
	Email string
}

// IdentityProvider resolves an authenticated session into a StreamForge identity.
type IdentityProvider interface {
	GetIdentity(
		ctx context.Context,
		sessionToken string,
	) (*Identity, error)
}
