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

// Registrar handles user registration through an authentication provider.
type Registrar interface {
	CreateRegistrationFlow(
		ctx context.Context,
	) (*RegistrationFlow, error)

	CompleteRegistration(
		ctx context.Context,
		input RegistrationInput,
	) (*RegistrationResult, error)
}

// RegistrationInput contains the information required to register a user.
type RegistrationInput struct {
	FlowID    string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// RegistrationResult represents the result of a successful registration.
type RegistrationResult struct {
	IdentityID string
	Email      string
}

// RegistrationFlow represents a user registration flow.
type RegistrationFlow struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}
