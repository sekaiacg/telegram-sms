package handler

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type StartHandler struct{}

func NewStartHandler() Handler {
	return &StartHandler{}
}

func (h *StartHandler) Command() string {
	return "start"
}

func (h *StartHandler) Description() string {
	return "Start the bot"
}

func (h *StartHandler) Handle(b *gotgbot.Bot, ctx *ext.Context) error {
	message := `
Hello, *%s %s*!
Thanks for using this bot.
Your UID is *%d*
	`
	_, err := ctx.EffectiveMessage.Reply(b,
		fmt.Sprintf(message, ctx.EffectiveUser.FirstName, ctx.EffectiveUser.LastName, ctx.EffectiveUser.Id),
		&gotgbot.SendMessageOpts{
			ParseMode: "markdown",
		},
	)
	return err
}
