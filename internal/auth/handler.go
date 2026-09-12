/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Package auth provides HTTP handlers for StreamForge authentication flows.

package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	platformauth "github.com/zuxtech/streamforge/internal/platform/auth"
	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
)

// Handler handles authentication-related HTTP requests.
type Handler struct {
	registrar platformauth.Registrar
	register platformauth.Registrar
}

// NewHandler creates an authentication HTTP handler.
func NewHandler(registrar platformauth.Registrar) *Handler {
	return &Handler{
		registrar: registrar,
	}
}

// CreateRegistrationFlow creates a new user registration flow.
func (h *Handler) CreateRegistrationFlow(
	w http.ResponseWriter,
	r *http.Request,
) {
	flow, err := h.registrar.CreateRegistrationFlow(r.Context())
	if err != nil {
		platformhttp.WriteJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to create registration flow",
			},
		)
		return
	}

	platformhttp.WriteJSON(
		w,
		http.StatusOK,
		flow,
	)
}


func (h *Handler) CompleteRegistration(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input platformauth.RegistrationInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
		return
	}

	if input.FlowID == "" {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "flow_id is required",
			},
		)
		return
	}

	if input.Email == "" {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "email is required",
			},
		)
		return
	}

	if input.Password == "" {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "password is required",
			},
		)
		return
	}


	result, err := h.registrar.CompleteRegistration(
		r.Context(),
		input,
	)


	if err != nil {
		if errors.Is(err, platformauth.ErrPasswordBreached) {
			platformhttp.WriteJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error":   "password_breached",
					"message": "The password has been found in data breaches and cannot be used.",
				},
			)
			return
		}

		platformhttp.WriteJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to complete registration",
			},
		)
		return
	}


	platformhttp.WriteJSON(
		w,
		http.StatusCreated,
		result,
	)
}
