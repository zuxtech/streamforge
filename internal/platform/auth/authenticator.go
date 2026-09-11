/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Auth defines the authentication contract used by StreamForge.

package auth

import "context"

type Identity struct {
	ID    string
	Email string
}

type Authenticator interface {
	GetIdentity(ctx context.Context, sessionToken string) (*Identity, error)
}