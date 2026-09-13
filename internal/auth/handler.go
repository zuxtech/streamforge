/*
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
	"net/http"

	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
)

// Handler handles authentication-related HTTP requests.
type Handler struct {
	registration *RegistrationService
}

// NewHandler creates an authentication HTTP handler.
func NewHandler(registration *RegistrationService) *Handler {
	return &Handler{
		registration: registration,
	}
}

type registrationRequest struct {
	FlowID    string `json:"flow_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// CreateRegistrationFlow creates a new user registration flow.
func (h *Handler) CreateRegistrationFlow(
	w http.ResponseWriter,
	r *http.Request,
) {
	flow, err := h.registration.CreateRegistrationFlow(
		r.Context(),
	)
	
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

// CompleteRegistration creates a StreamForge user account.
func (h *Handler) CompleteRegistration(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req registrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
		return
	}

	if req.FlowID == "" {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "flow_id is required",
			},
		)
		return
	}

	if req.Email == "" {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "email is required",
			},
		)
		return
	}

	if req.Password == "" {
		platformhttp.WriteJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "password is required",
			},
		)
		return
	}

	input := RegistrationInput{
		FlowID:    req.FlowID,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}


	result, err := h.registration.Register(
		r.Context(),
		input,
	)
	
	if err != nil {
		platformhttp.WriteError(w, err)
		return
	}

	platformhttp.WriteJSON(
		w,
		http.StatusCreated,
		result,
	)
}

func (h *Handler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {
	identity, ok := IdentityFromContext(r.Context())

	if !ok {
		platformhttp.WriteJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "unauthorized",
			},
		)
		return
	}

	platformhttp.WriteJSON(
		w,
		http.StatusOK,
		identity,
	)
}
