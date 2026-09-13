 /*
  * Copyright (c) 2026 ZuxTech.
  *
  * SPDX-License-Identifier: Apache-2.0
  *
  * This source code is licensed under the Apache License, Version 2.0
  * found in the LICENSE file in the root directory of this source tree.
  */

 package kratos

 import (
 	"fmt"
 	"strings"

 	"github.com/zuxtech/streamforge/internal/platform/errors"
 )

 // mapRegistrationError maps an ORY Kratos SDK error to a StreamForge
 // registration error.
 func mapRegistrationError(err error) error {
 	if err == nil {
 		return nil
 	}

 	message := strings.ToLower(err.Error())

 	if isPasswordBreached(message) {
 		return errors.ErrPasswordBreached
 	}

 	if strings.Contains(message, "expired") {
 		return errors.ErrRegistrationFlowExpired
 	}

 	if isUserAlreadyExists(message) {
 		return errors.ErrUserAlreadyExists
 	}

 	return fmt.Errorf(
 		"%w: %v",
 		errors.ErrRegistrationFailed,
 		err,
 	)
 }

 // isUserAlreadyExists reports whether the Kratos error indicates
 // that the identity already exists.
 func isUserAlreadyExists(message string) bool {
 	return strings.Contains(message, "already exists") ||
 		strings.Contains(message, "already registered") ||
 		strings.Contains(message, "identity with") &&
 			strings.Contains(message, "already")
 }

 // isPasswordBreached reports whether the Kratos error indicates
 // that the supplied password has appeared in a known data breach.
 func isPasswordBreached(message string) bool {
 	return strings.Contains(message, "data breach") ||
 		strings.Contains(message, "breached password") ||
 		strings.Contains(message, "password has been found")
 }

 // mapSessionError maps an ORY Kratos SDK error to a StreamForge
 // session error.
 func mapSessionError(err error) error {
 	if err == nil {
 		return nil
 	}

 	message := strings.ToLower(err.Error())

 	if strings.Contains(message, "unauthorized") ||
 		strings.Contains(message, "invalid session") ||
 		strings.Contains(message, "session is invalid") {
 		return errors.ErrInvalidSession
 	}

 	return fmt.Errorf(
 		"kratos whoami request failed: %w",
 		err,
 	)
 }