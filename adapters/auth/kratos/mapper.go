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
	ory "github.com/ory/kratos-client-go/v26"

	"github.com/zuxtech/streamforge/internal/auth"
)

// mapIdentity maps an Ory Kratos identity to a StreamForge identity.
func mapIdentity(response ory.Identity) *auth.Identity {
	return &auth.Identity{
		ID:    response.GetId(),
		Email: getIdentityEmail(response),
	}
}

// mapRegistrationFlow maps an Ory Kratos registration flow to a
// StreamForge registration flow.
func mapRegistrationFlow(
	response *ory.RegistrationFlow,
) *auth.RegistrationFlow {
	return &auth.RegistrationFlow{
		ID:     response.Id,
		Action: response.Ui.Action,
	}
}

// mapRegistrationResult maps a successful Ory Kratos registration
// response to a StreamForge registration result.
func mapRegistrationResult(
	response *ory.SuccessfulNativeRegistration,
) *auth.RegistrationResult {
	identity := response.GetIdentity()

	return &auth.RegistrationResult{
		IdentityID: identity.GetId(),
		Email:      getIdentityEmail(identity),
	}
}

// getIdentityEmail extracts the email address from an Ory Kratos identity.
func getIdentityEmail(identity ory.Identity) string {
	traits, ok := identity.GetTraits().(map[string]interface{})
	if !ok {
		return ""
	}

	email, ok := traits["email"].(string)
	if !ok {
		return ""
	}

	return email
}