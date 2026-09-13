/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package config

type Config struct {
	Server      ServerConfig      `yaml:"server"`
	Auth        AuthConfig        `yaml:"auth"`
	Database    DatabaseConfig    `yaml:"database"`
	Redis       RedisConfig       `yaml:"redis"`
	Storage     StorageConfig     `yaml:"storage"`
	Transcoding TranscodingConfig `yaml:"transcoding"`
}

type ServerConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	PoweredBy string `yaml:"powered_by"`

	Addr    string `yaml:"-"`
	BaseURL string `yaml:"-"`
}

type AuthConfig struct {
	Provider string       `yaml:"provider"`
	Kratos   KratosConfig `yaml:"kratos"`
}

type KratosConfig struct {
	URL string `yaml:"url"`
}

type DatabaseConfig struct {
	URL string `yaml:"url"`
}

type RedisConfig struct {
	URL string `yaml:"url"`
}

type StorageConfig struct {
	Provider string `yaml:"provider"`
}

type TranscodingConfig struct {
	Provider string `yaml:"provider"`
}
