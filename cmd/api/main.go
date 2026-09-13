/*
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/zuxtech/streamforge/adapters/auth/kratos"
	"github.com/zuxtech/streamforge/internal/auth"
	"github.com/zuxtech/streamforge/internal/platform/config"
	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
	"github.com/zuxtech/streamforge/internal/user"
)

func main() {
	cfg := config.Load()

	fmt.Println(cfg)

	server := platformhttp.NewServer(
	platformhttp.WithPoweredBy(cfg.Server.PoweredBy),
	)

	r := server.Router()

	kratosClient := kratos.NewClient(
		cfg.Auth.Kratos.URL,
	)

	identityProvider := kratos.NewAuthenticator(
		kratosClient,
	)

	registrationAdapter := kratos.NewRegistrationAdapter(
		kratosClient,
	)


	pool, err := pgxpool.New(
		context.Background(),
		cfg.Database.URL,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepository := user.NewPostgresRepository(pool)

	registrationService := auth.NewRegistrationService(
		registrationAdapter,
		userRepository,
	)

	authHandler := auth.NewHandler(
		registrationService,
	)

	server.RegisterHealthRoutes(r)

	r.Route("/v1", func(r chi.Router) {
		authHandler.RegisterRoutes(r, identityProvider)
	})

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
