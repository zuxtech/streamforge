/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package user

import "context"

type Repository interface {
	Create(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByKratosIdentityID(ctx context.Context, kratosIdentityID string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
