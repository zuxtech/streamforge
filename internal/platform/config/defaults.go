/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package config

const (
	defaultConfigFile = "streamforge.yml"

	defaultHost      = "127.0.0.1"
	defaultPort      = 8080
	defaultPoweredBy = "StreamForge"

	defaultKratosURL   = "http://127.0.0.1:4433"
	defaultDatabaseURL = "postgres://localhost/streamforge"
	defaultRedisURL    = "redis://127.0.0.1:6379"

	defaultAuthProvider        = "kratos"
	defaultStorageProvider     = "s3"
	defaultTranscodingProvider = "ffmpeg"
) 

func defaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Host:      defaultHost,
			Port:      defaultPort,
			PoweredBy: defaultPoweredBy,
		},
		Auth: AuthConfig{
			Provider: defaultAuthProvider,
			Kratos: KratosConfig{
				URL: defaultKratosURL,
			},
		},
		Database: DatabaseConfig{
			URL: defaultDatabaseURL,
		},
		Redis: RedisConfig{
			URL: defaultRedisURL,
		},
		Storage: StorageConfig{
			Provider: defaultStorageProvider,
		},
		Transcoding: TranscodingConfig{
			Provider: defaultTranscodingProvider,
		},
	}
}
