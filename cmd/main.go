package main

import (
	"fmt"
	"os"

	"github.com/far4599/swagger-openapiv2-merge/internal/app"
)

var (
	version = "dev"
)

func main() {
	if err := app.NewApp(version).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
