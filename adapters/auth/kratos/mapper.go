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

import "github.com/zuxtech/streamforge/internal/platform/auth"

func mapIdentity(response *whoAmIResponse) *auth.Identity {
	return &auth.Identity{
		ID:    response.ID,
		Email: response.Traits.Email,
	}
}