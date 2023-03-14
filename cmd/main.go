package main

import (
	"api-go-example/cmd/config"
	"fmt"
)

func main() {
	cfg := config.GetAPIConfig()

	fmt.Printf("configs %v\n", cfg)
	// init logger

	// init client, service, endpoint, handler
}
