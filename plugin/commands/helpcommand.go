package commands

import (
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/container"
	"github.com/Syfaro/telegram-bot-api"
)

type HelpCommand struct {
	hasCompleted bool
	logger *logging.Logger
	originalMessage *tgbotapi.Message
}

func CreateHelpCommand(logger *logging.Logger) *HelpCommand {
	return &HelpCommand{
		hasCompleted: false,
		logger: logger,
	}
}

func (cmd *HelpCommand) HasCompleted() bool {
	return cmd.hasCompleted
}

func (cmd *HelpCommand) SetRequestMessage(message *tgbotapi.Message) {
	cmd.originalMessage = message
}

func (cmd *HelpCommand) GetResponseMessage() (*container.Response, error) {
	config := tgbotapi.NewMessage(cmd.originalMessage.Chat.ID, cmd.originalMessage.Text)
	config.Text = "TODO: Help message."

	response := &container.Response{
		ResponseConfig: config,
		ConfigType: "MessageConfig",
	}

	cmd.hasCompleted = true

	return response, nil
}
