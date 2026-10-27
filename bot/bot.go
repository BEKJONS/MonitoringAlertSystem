package bot

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// ConnectToTelegram creates the bot from the TELEGRAM_BOT_TOKEN environment variable.
func ConnectToTelegram() (*telego.Bot, error) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, errors.New("TELEGRAM_BOT_TOKEN is not set")
	}

	bot, err := telego.NewBot(token)
	if err != nil {
		return nil, err
	}
	// Call method getMe (https://core.telegram.org/bots/api#getme)
	botUser, err := bot.GetMe()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Bot user: %s\n", botUser.Username)

	return bot, nil
}

func SendMessageToAllUsers(bot *telego.Bot, message string) error {
	// Get the list of all chat IDs
	updates, err := bot.GetUpdates(nil)
	if err != nil {
		return err
	}

	// Send the message to each user
	sent := 0
	for _, update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		_, err := bot.SendMessage(tu.Message(tu.ID(chatID), message))
		if err != nil {
			log.Printf("Failed to send message to user %d: %v", chatID, err)
			continue
		}
		sent++
	}

	fmt.Printf("Message sent to %d users\n", sent)
	return nil
}
