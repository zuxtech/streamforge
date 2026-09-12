/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func applyEnvironment(cfg *Config) {
	cfg.Server.Host = getStringEnv(
		"STREAMFORGE_API_HOST",
		cfg.Server.Host,
	)

	cfg.Server.Port = getIntEnv(
		"STREAMFORGE_API_PORT",
		cfg.Server.Port,
	)

	cfg.Auth.Provider = getStringEnv(
		"STREAMFORGE_AUTH_PROVIDER",
		cfg.Auth.Provider,
	)

	cfg.Auth.Kratos.URL = getStringEnv(
		"STREAMFORGE_KRATOS_URL",
		cfg.Auth.Kratos.URL,
	)

	cfg.Database.URL = getStringEnv(
		"STREAMFORGE_DATABASE_URL",
		cfg.Database.URL,
	)

	cfg.Redis.URL = getStringEnv(
		"STREAMFORGE_REDIS_URL",
		cfg.Redis.URL,
	)

	cfg.Storage.Provider = getStringEnv(
		"STREAMFORGE_STORAGE_PROVIDER",
		cfg.Storage.Provider,
	)

	cfg.Transcoding.Provider = getStringEnv(
		"STREAMFORGE_TRANSCODING_PROVIDER",
		cfg.Transcoding.Provider,
	)

	if value, ok := getOptionalStringEnv("STREAMFORGE_API_POWERED_BY"); ok {
		cfg.Server.PoweredBy = value
	}
}

func getStringEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return fallback
	}

	return value
}

func getIntEnv(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return fallback
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		log.Printf(
			"warning: invalid %s=%q, using %d",
			key,
			value,
			fallback,
		)

		return fallback
	}

	return result
}

func getOptionalStringEnv(key string) (string, bool) {
	value, exists := os.LookupEnv(key)

	if !exists {
		return "", false
	}

	return strings.TrimSpace(value), true
}
