package telegram

import (
	"errors"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Client) setupCommands() error {
	commands := []models.BotCommand{
		{
			Command:     "/start",
			Description: "Hey there! I can save you videos. Just send me a link",
		},
		{
			Command:     "/help",
			Description: "I can help you to save videos to local store",
		},
		{
			Command:     "/download",
			Description: "Download video by link",
		},
	}

	success, err := c.bot.SetMyCommands(c.ctx, &bot.SetMyCommandsParams{
		Commands:     commands,
		Scope:        nil,
		LanguageCode: "",
	})
	if err != nil {
		return err
	}
	if !success {
		return errors.New("setMyCommands failed")
	}
	return nil
}
