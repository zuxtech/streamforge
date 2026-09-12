/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package auth defines authentication and registration contracts used by StreamForge.
package auth

import "context"

// RegistrationFlow represents a provider-managed user registration flow.


type RegistrationFlow struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}

type RegistrationInput struct {
    FlowID    string `json:"flow_id"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

type RegistrationResult struct {
    IdentityID string
    Email      string
}

// Registrar creates registration flows for user onboarding.
type Registrar interface {
	CreateRegistrationFlow(ctx context.Context) (*RegistrationFlow, error)
	 CompleteRegistration(
        ctx context.Context,
        input RegistrationInput,
    ) (*RegistrationResult, error)
}
