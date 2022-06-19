package main

import (
	"fmt"
	"os"

	"github.com/far4599/swagger-openapiv2-merge/internal/app"
)

func main() {
	if err := app.NewApp().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
