/*
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package http provides HTTP infrastructure for StreamForge.
package http

import (
	"encoding/json"
	"errors"
	"net/http"

	sferrors "github.com/zuxtech/streamforge/internal/platform/errors"
)

// ErrorResponse represents the public StreamForge API error format.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// WriteError writes a consistent StreamForge API error response.
func WriteError(
	w http.ResponseWriter,
	err error,
) {
	status, response := mapError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(response)
}

// mapError maps a StreamForge application error to a public HTTP error.
func mapError(err error) (int, ErrorResponse) {
	switch {
	case errors.Is(err, sferrors.ErrInvalidSession):
		return http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_session",
			Message: "Your session is invalid or has expired.",
		}


	case errors.Is(err, sferrors.ErrUserAlreadyExists):
        return http.StatusConflict, ErrorResponse{
            Error:   "user_already_exists",
            Message: "a user with this email already exists",
        }

	case errors.Is(err, sferrors.ErrPasswordBreached):
		return http.StatusBadRequest, ErrorResponse{
			Error:   "password_breached",
			Message: "The password has been found in data breaches and cannot be used.",
		}

	case errors.Is(err, sferrors.ErrRegistrationFlowExpired):
		return http.StatusGone, ErrorResponse{
			Error:   "registration_flow_expired",
			Message: "The registration flow has expired.",
		}

	case errors.Is(err, sferrors.ErrRegistrationFailed):
		return http.StatusBadRequest, ErrorResponse{
			Error:   "registration_failed",
			Message: "Registration could not be completed.",
		}

	default:
		return http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "An unexpected error occurred.",
		}
	}
}
