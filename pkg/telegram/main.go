package telegram

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tashima42/currency-watcher/pkg/currencyconverter"
	"github.com/tashima42/currency-watcher/pkg/currencyprovider"
	"go.uber.org/zap"
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

func (t *Telegram) HandleUpdates(logger *zap.Logger) {
	logger.Info("Getting updates from telegram")
	updates := t.bot.GetUpdatesChan(t.updateConfig)
	currencySetTo := currencyprovider.CLP

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		switch update.Message.Command() {
		case "help":
			t.SendMessage(update.Message.Chat.ID, "Commands:\n/help: see this message\n/status: check bot status\n/clp: convert CLP to BRL\n/brl: convert BRL to CLP")
		case "status":
			t.SendMessage(update.Message.Chat.ID, "(┛ಠ_ಠ)┛彡┻━┻")
		case "clp":
			currencySetTo = currencyprovider.CLP
			t.SendMessage(update.Message.Chat.ID, "Set currency to CLP")
		case "brl":
			currencySetTo = currencyprovider.BRL
			t.SendMessage(update.Message.Chat.ID, "Set currency to BRL")
		case "currency":
			t.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Current currency is: %s", currencySetTo))
		default:
			fmt.Println("default")
			amount, err := strconv.ParseFloat(strings.ReplaceAll(update.Message.Text, "/", ""), 64)
			if err != nil {
				logger.Error("Failed to parse amount string to float: " + err.Error())
				t.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Failed to convert, error message: %s", err.Error()))
				continue
			}
			to := currencyprovider.BRL
			if currencySetTo == currencyprovider.BRL {
				to = currencyprovider.CLP
			}
			msg, err := currencyconverter.Convert(currencySetTo, to, amount)
			if err != nil {
				logger.Error("Failed to convert currency: " + err.Error())
				t.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Failed to convert, error message: %s", err.Error()))
				continue
			}
			t.SendMessage(update.Message.Chat.ID, msg)
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
