package main

import (
	"fmt"

	"github.com/pablotrianda/monkey-clocker/api"
)

func main() {
	fmt.Println("🐒⏰ MONKEY-CLOCKER")
	api.Start("9090")
}
