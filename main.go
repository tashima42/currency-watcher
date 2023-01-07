package main

import (
	"fmt"
	"os"

	"github.com/tashima42/currency-watcher/cmd/currencywatcher"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Info("Starting command")
	rootCmd := currencywatcher.InitCommand(logger)
	logger.Info("Executing command")
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("Failed to execute command: " + err.Error())
	}
}
