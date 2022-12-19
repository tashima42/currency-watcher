package currencywatcher

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/tashima42/currency-watcher/pkg/currencywatcher"
	"github.com/tashima42/currency-watcher/pkg/telegram"
)

var (
	chatID            int64
	currencyThreshold float64
)

func InitCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "currencywatcher",
		Short: "Currency Watcher helps you to know when you can exchange currency",
		Long:  "Currency Watcher helps you to know when you can exchange currency",
		RunE: func(cmd *cobra.Command, args []string) error {
			telegram, err := telegram.NewBot(false)
			if err != nil {
				fmt.Println(err)
				return err
			}
			telegram.ConfigBot()

			c := cron.New()
			c.AddFunc("35 8-21 * * *", func() {
				message, err := currencywatcher.Check(currencyThreshold, false)
				if err != nil {
					fmt.Println(err)
				}
				if message != nil {
					telegram.SendMessage(chatID, *message)
				}
			})

			c.Start()
			telegram.HandleUpdates(chatID, currencyThreshold)
			return nil
		},
	}

	rootCmd.Flags().Int64VarP(&chatID, "chatID", "c", 1, "Default Chat ID")
	rootCmd.Flags().Float64VarP(&currencyThreshold, "currency-threshold", "t", 5.25, "Currency Threshold to notify user")

	return rootCmd
}
