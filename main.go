package main

import (
	"fmt"
	"os"

	"github.com/tashima42/currency-watcher/cmd/currencywatcher"
)

func main() {
	rootCmd := currencywatcher.InitCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
