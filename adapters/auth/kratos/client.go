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
	"net/http"
	"strings"
	"time"

	ory "github.com/ory/kratos-client-go/v26"
)

type Client struct {
	api *ory.APIClient
}

func NewClient(baseURL string) *Client {
	config := ory.NewConfiguration()

	config.Servers = ory.ServerConfigurations{
		{
			URL: strings.TrimRight(baseURL, "/"),
		},
	}

	config.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		api: ory.NewAPIClient(config),
	}
}
