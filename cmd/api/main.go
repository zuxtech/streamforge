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

	"github.com/zuxtech/streamforge/adapters/auth/kratos"
	"github.com/zuxtech/streamforge/internal/auth"
	"github.com/zuxtech/streamforge/internal/platform/config"
	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
	"github.com/zuxtech/streamforge/internal/user"
)

func main() {
	cfg := config.Load()

	kratosClient := kratos.NewClient(cfg.Auth.Kratos.URL)

	authenticator := kratos.NewAuthenticator(kratosClient)
	registrar := kratos.NewRegistrar(kratosClient)

	server := platformhttp.NewServer(
		platformhttp.WithPoweredBy(cfg.Server.PoweredBy),
	)

	userHandler := &user.Handler{}

	userHandler.RegisterRoutes(
		server.Mux(),
		authenticator,
	)

	authHandler := auth.NewHandler(registrar)

	authHandler.RegisterRoutes(
		server.Mux(),
	)

	log.Printf(
		"streamforge api listening on %s",
		cfg.Server.BaseURL,
	)

	if err := http.ListenAndServe(
		cfg.Server.Addr,
		server,
	); err != nil {
		log.Fatal(err)
	}
}
