package main

import (
	"url-shortener/cmd/internal/config"
)

func main() {

	cfg := config.MustLoad()

	// TODO: init logger: slog

	// TODO: init storage: postgresql

	// TODO: init router: chi, chi render

	// TODO: run server
}
