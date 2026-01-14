package main

import (
	"fmt"

	config "github.com/Adopten123/go-urlshortener/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println(config)
	// TODO: init logger. Libs: log/slog
	// TODO: init storage. Libs: sqlite
	// TODO: init router. Libs: chi, "chi render"
	// TODO: run server.
}
