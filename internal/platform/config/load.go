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
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() Config {
	cfg := defaultConfig()

	defaultScheme := "http"

	loadDotEnv()

	configFile := getStringEnv(
		"STREAMFORGE_CONFIG_FILE",
		defaultConfigFile,
	)

	if err := loadYAML(configFile, &cfg); err != nil {
		log.Fatalf("failed to load config file %q: %v", configFile, err)
	}

	applyEnvironment(&cfg)

	addr := net.JoinHostPort(
		cfg.Server.Host,
		strconv.Itoa(cfg.Server.Port),
	)

	cfg.Server.Addr = addr
	cfg.Server.BaseURL = defaultScheme + "://" + addr

	return cfg
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			log.Printf("warning: failed to load .env: %v", err)
		}
	}
}
