package telegram

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tashima42/currency-watcher/pkg/currencywatcher"
)

type Telegram struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
}

func NewBot(debug bool) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		return nil, err
	}
	bot.Debug = debug
	t := Telegram{}
	t.bot = bot
	return &t, err
}

func (t *Telegram) ConfigBot() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
}

func (t *Telegram) HandleUpdates(chatID int64, currencyThreshold float64) {
	updates := t.bot.GetUpdatesChan(t.updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		switch update.Message.Command() {
		case "help":
			t.SendMessage(update.Message.Chat.ID, "I understand /check and /status.")
		case "status":
			t.SendMessage(update.Message.Chat.ID, "(┛ಠ_ಠ)┛彡┻━┻")
		case "check":
			msg, err := currencywatcher.Check(currencyThreshold, true)
			if err != nil {
				fmt.Println(err)
				return
			}
			t.SendMessage(chatID, *msg)
		default:
			t.SendMessage(update.Message.Chat.ID, "I don't know that command")
		}

	}
}

func (t *Telegram) SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = text
	if _, err := t.bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}
