package config

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	APIAddr string
	BaseURL string
}

const (
	defaultHost = "127.0.0.1"
	defaultPort = 8080
	defaultScheme = "http"
)

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Print("warning: .env not loaded: ", err)
	}

	port := getPort("STREAMFORGE_API_PORT", defaultPort)
	host := getHost("STREAMFORGE_API_HOST", defaultHost)

	addr := net.JoinHostPort(host, strconv.Itoa(port))

	baseURL := defaultScheme + "://" + addr

	return Config{
		APIAddr: addr,
		BaseURL: baseURL,
	}
}

func getPort(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	port, err := strconv.Atoi(value)
	if err != nil {
		log.Printf(
			"warning: invalid %s=%q, using %d",
			key,
			value,
			fallback,
		)

		return fallback
	}

	return port
}

func getHost(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	host := strings.TrimSpace(value)

	if host == "" {
		log.Printf(
			"warning: invalid %s=%q, using %s",
			key,
			value,
			fallback,
		)

		return fallback
	}

	return host
}