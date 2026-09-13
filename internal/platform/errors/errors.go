/*
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package errors provides common StreamForge application errors.
package errors

import "errors"

var (
	ErrInvalidSession   = errors.New("invalid session")
	ErrPasswordBreached = errors.New(
		"password has been found in data breaches",
	)
	ErrRegistrationFlowExpired = errors.New(
		"registration flow expired",
	)

	ErrRegistrationFailed = errors.New(
		"registration failed",
	)
	  ErrUserAlreadyExists       = errors.New("user already exists")
)
