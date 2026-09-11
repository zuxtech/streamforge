/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package main

import (
	"log"
	"net/http"

	"github.com/zuxtech/streamforge/internal/platform/config"
	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
)

func main() {
	cfg := config.Load()

	handler := platformhttp.NewServer()

	addr := cfg.APIAddr
	baseURL := cfg.BaseURL

	if addr == "" {
		addr = ":8080"
	}

	log.Printf("streamforge api listening on %s", baseURL)

	if err := http.ListenAndServe(addr, handler.Handler()); err != nil {
		log.Fatal(err)
	}
}