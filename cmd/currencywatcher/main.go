package currencywatcher

import (
	"github.com/spf13/cobra"
	"github.com/tashima42/currency-watcher/pkg/telegram"
	"go.uber.org/zap"
)

var debug bool

func InitCommand(logger *zap.Logger) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "currencywatcher",
		Short: "Currency Watcher helps you to know when you can exchange currency",
		Long:  "Currency Watcher helps you to know when you can exchange currency",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info("Starting telegram bot")
			telegram, err := telegram.NewBot(debug)
			if err != nil {
				logger.Error("Failed to start telegram bot" + err.Error())
				return err
			}
			logger.Info("Configuring telegram bot")
			telegram.ConfigBot()

			logger.Info("Starting to handle updates in telegram bot")
			telegram.HandleUpdates(logger)
			return nil
		},
	}

	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Telegram debug mode")

	return rootCmd
}
