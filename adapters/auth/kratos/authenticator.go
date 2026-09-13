/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Authenticator adapts ORY Kratos authentication to the StreamForge auth interface.

package kratos

import (
	"context"

	"github.com/zuxtech/streamforge/internal/auth"
)

type Authenticator struct {
	client *Client
}


var _ auth.IdentityProvider = (*Authenticator)(nil)

func NewAuthenticator(client *Client) *Authenticator {
	return &Authenticator{
		client: client,
	}
}

func (a *Authenticator) GetIdentity(
	ctx context.Context,
	sessionToken string,
) (*auth.Identity, error) {
	response, _, err := a.client.api.FrontendAPI.
		ToSession(ctx).
		XSessionToken(sessionToken).
		Execute()

	if err != nil {
		return nil, mapSessionError(err)
	}

	return mapIdentity(response.GetIdentity()), nil
}