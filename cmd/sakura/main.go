package main

import (
	"fmt"

	"github.com/s-chernyavskiy/sakura/internal/sakura/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
}
