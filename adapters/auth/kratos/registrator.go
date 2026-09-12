/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package kratos provides the StreamForge authentication adapter for Ory Kratos.
package kratos

import (
	"context"

	"github.com/zuxtech/streamforge/internal/platform/auth"
)

// Registrar implements user registration through Ory Kratos.
type Registrar struct {
	client *Client
}

// NewRegistrar creates a Kratos registration adapter.
func NewRegistrar(client *Client) *Registrar {
	return &Registrar{
		client: client,
	}
}

func (r *Registrar) CreateRegistrationFlow(
	ctx context.Context,
) (*auth.RegistrationFlow, error) {
	response, err := r.client.createRegistrationFlow(ctx)
	if err != nil {
		return nil, err
	}

	return mapRegistrationFlow(response), nil
}


func (r *Registrar) CompleteRegistration(
    ctx context.Context,
    input auth.RegistrationInput,
) (*auth.RegistrationResult, error) {
    response, err := r.client.completeRegistration(ctx, input)
    if err != nil {
        return nil, err
    }

    return mapRegistrationResult(response), nil
}