/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

// Client isolates communication between ORY Kratos and StreamForge.

package kratos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

type whoAmIResponse struct {
	ID         string `json:"id"`
	Traits     struct {
		Email string `json:"email"`
	} `json:"traits"`
}

func (c *Client) whoAmI(
	ctx context.Context,
	sessionToken string,
) (*whoAmIResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/sessions/whoami",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create kratos request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+sessionToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call kratos: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"kratos returned status %d",
			resp.StatusCode,
		)
	}

	var result whoAmIResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode kratos response: %w", err)
	}

	return &result, nil
}