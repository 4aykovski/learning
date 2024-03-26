package main

import (
	"123/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config")
	}

	fmt.Println(cfg.STORAGE)

}
