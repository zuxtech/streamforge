/*
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package auth

import (
	"context"

	"github.com/zuxtech/streamforge/internal/platform/id"
	"github.com/zuxtech/streamforge/internal/user"
)

// RegistrationService coordinates user registration.
type RegistrationService struct {
	registrar Registrar
	users     user.Repository
}

// NewRegistrationService creates a registration service.
func NewRegistrationService(
	registrar Registrar,
	users user.Repository,
) *RegistrationService {
	return &RegistrationService{
		registrar: registrar,
		users:     users,
	}
}

// CreateRegistrationFlow creates a new registration flow.
func (s *RegistrationService) CreateRegistrationFlow(
	ctx context.Context,
) (*RegistrationFlow, error) {
	return s.registrar.CreateRegistrationFlow(ctx)
}

// Register creates a StreamForge user through the authentication provider.
func (s *RegistrationService) Register(
	ctx context.Context,
	input RegistrationInput,
) (*user.User, error) {
	result, err := s.registrar.CompleteRegistration(
		ctx,
		input,
	)
	
	if err != nil {
		return nil, err
	}

	u := &user.User{
		ID:               id.New("usr"),
		KratosIdentityID: result.IdentityID,
		Email:            result.Email,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		EmailVerified:    false,
	}


	if err := s.users.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
