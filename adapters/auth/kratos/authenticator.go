/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Authenticator adapts ORY Kratos authentication to StreamForge.
package kratos

import (
	"context"

	"github.com/zuxtech/streamforge/internal/platform/auth"
)

type Authenticator struct {
	client *Client
}

func NewAuthenticator(client *Client) *Authenticator {
	return &Authenticator{
		client: client,
	}
}

func (a *Authenticator) GetIdentity(
	ctx context.Context,
	sessionToken string,
) (*auth.Identity, error) {
	response, err := a.client.whoAmI(ctx, sessionToken)
	if err != nil {
		return nil, err
	}

	return mapIdentity(response), nil
}