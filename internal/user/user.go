/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package user

import "time"

type User struct {
	ID               string    `json:"id"`
	KratosIdentityID string    `json:"kratosIdentityId"`
	Email            string    `json:"email"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	EmailVerified    bool      `json:"emailVerified"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}