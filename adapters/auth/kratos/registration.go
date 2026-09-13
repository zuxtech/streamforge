/*
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

	ory "github.com/ory/kratos-client-go/v26"
	"github.com/zuxtech/streamforge/internal/auth"
)

type RegistrationAdapter struct {
	client *Client
}


// RegistrationAdapter implements user registration through Ory Kratos.

var _ auth.Registrar = (*RegistrationAdapter)(nil)

func NewRegistrationAdapter(client *Client) *RegistrationAdapter {
	return &RegistrationAdapter{
		client: client,
	}
}

// CreateRegistrationFlow creates a new registration flow.
func (r *RegistrationAdapter) CreateRegistrationFlow(
	ctx context.Context,
) (*auth.RegistrationFlow, error) {
		response, _, err := r.client.api.FrontendAPI.
		CreateNativeRegistrationFlow(ctx).
		Execute()


	if err != nil {
		return nil, err
	}

	return mapRegistrationFlow(response), nil
}

// CompleteRegistration completes a registration flow.
func (r *RegistrationAdapter) CompleteRegistration(
	ctx context.Context,
	input auth.RegistrationInput,
) (*auth.RegistrationResult, error) {
	payload := ory.NewUpdateRegistrationFlowWithPasswordMethod(
		"password",
		input.Password,
		map[string]interface{}{
			"email": input.Email,
			"name": map[string]interface{}{
				"first": input.FirstName,
				"last":  input.LastName,
			},
		},
	)

	body := ory.UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody(
		payload,
	)

	response, _, err := r.client.api.FrontendAPI.
		UpdateRegistrationFlow(ctx).
		Flow(input.FlowID).
		UpdateRegistrationFlowBody(body).
		Execute()

	if err != nil {
		return nil, mapRegistrationError(err)
	}

	return mapRegistrationResult(response), nil
}