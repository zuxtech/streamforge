/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Mapper translates ORY Kratos responses into StreamForge authentication types.

package kratos

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zuxtech/streamforge/internal/platform/auth"
)

func mapIdentity(response *whoAmIResponse) *auth.Identity {
	return &auth.Identity{
		ID:    response.ID,
		Email: response.Traits.Email,
	}
}

func mapRegistrationFlow(
	response *registrationFlowResponse,
) *auth.RegistrationFlow {
	return &auth.RegistrationFlow{
		ID:     response.ID,
		Action: response.UI.Action,
	}
}

func mapRegistrationResult(
    response *registrationResponse,
) *auth.RegistrationResult {
    return &auth.RegistrationResult{
        IdentityID: response.Identity.ID,
        Email:      response.Identity.Traits.Email,
    }
}


func mapRegistrationError(
	resp *http.Response,
) error {
	var result kratosErrorResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf(
			"kratos registration endpoint returned status %d",
			resp.StatusCode,
		)
	}

	for _, node := range result.UI.Nodes {
		for _, message := range node.Messages {
			if message.Text == "The password has been found in data breaches and must no longer be used." {
				return auth.ErrPasswordBreached
			}
		}
	}

	return fmt.Errorf(
		"kratos registration endpoint returned status %d",
		resp.StatusCode,
	)
}