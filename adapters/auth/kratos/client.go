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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/zuxtech/streamforge/internal/platform/auth"
)

// Client communicates with the ORY Kratos public API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new ORY Kratos API client.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{},
	}
}

type Traits struct {
    Email string `json:"email"`
}

type whoAmIResponse struct {
	ID     string `json:"id"`
	Traits Traits `json:"traits"`
}

// registrationFlowResponse represents the registration flow returned by Kratos.
type registrationFlowResponse struct {
	ID     string `json:"id"`
	UI struct {
		Action string `json:"action"`
	} `json:"ui"`
}


type registrationRequest struct {
    Method   string `json:"method"`
    Password string `json:"password"`
    Traits   struct {
        Email string `json:"email"`
        Name  struct {
            First string `json:"first"`
            Last  string `json:"last"`
        } `json:"name"`
    } `json:"traits"`
}

type registrationResponse struct {
    Identity struct {
        ID     string `json:"id"`
        Traits struct {
            Email string `json:"email"`
        } `json:"traits"`
    } `json:"identity"`
}

type kratosMessage struct {
	ID      int                    `json:"id"`
	Text    string                 `json:"text"`
	Type    string                 `json:"type"`
	Context map[string]interface{} `json:"context"`
}

type kratosNode struct {
	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`

	Messages []kratosMessage `json:"messages"`
}

type kratosErrorResponse struct {
	UI struct {
		Nodes []kratosNode `json:"nodes"`
	} `json:"ui"`
}

// createRegistrationFlow creates a Kratos self-service registration flow.
func (c *Client) createRegistrationFlow(
	ctx context.Context,
) (*registrationFlowResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/self-service/registration/api",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create kratos registration request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call kratos registration endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"kratos registration endpoint returned status %d",
			resp.StatusCode,
		)
	}

	var result registrationFlowResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"decode kratos registration response: %w",
			err,
		)
	}

	return &result, nil
}


func (c *Client) completeRegistration(
    ctx context.Context,
    input auth.RegistrationInput,
) (*registrationResponse, error) {

    payload := registrationRequest{
        Method:   "password",
        Password: input.Password, 
    }

    payload.Traits.Email = input.Email
    payload.Traits.Name.First = input.FirstName
    payload.Traits.Name.Last = input.LastName

    body, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf(
            "encode kratos registration request: %w",
            err,
        )
    }

    req, err := http.NewRequestWithContext(
        ctx,
        http.MethodPost,
        c.baseURL+"/self-service/registration?flow="+input.FlowID,
        bytes.NewReader(body),
    )
    if err != nil {
        return nil, fmt.Errorf(
            "create kratos registration request: %w",
            err,
        )
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf(
            "call kratos registration endpoint: %w",
            err,
        )
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, mapRegistrationError(resp)
	}

    var result registrationResponse

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf(
            "decode kratos registration response: %w",
            err,
        )
    }

    return &result, nil
}



// whoAmI retrieves the identity associated with a Kratos session token.
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
		return nil, fmt.Errorf(
			"decode kratos response: %w",
			err,
		)
	}

	return &result, nil
}
