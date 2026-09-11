package main

import (
	"log"
	"net/http"

	"github.com/zuxtech/streamforge/internal/platform/config"
	platformhttp "github.com/zuxtech/streamforge/internal/platform/http"
)

func main() {
	cfg := config.Load()

	handler := platformhttp.HttpServer()

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