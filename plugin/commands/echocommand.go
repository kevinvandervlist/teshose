package commands

import (
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/container"
	"github.com/Syfaro/telegram-bot-api"
)

type EchoCommand struct {
	hasCompleted bool
	logger *logging.Logger
	originalMessage *tgbotapi.Message
}

func CreateEchoCommand(logger *logging.Logger) *EchoCommand {
	return &EchoCommand{
		hasCompleted: false,
		logger: logger,
	}
}

func (cmd *EchoCommand) HasCompleted() bool {
	return cmd.hasCompleted
}

func (cmd *EchoCommand) SetRequestMessage(message *tgbotapi.Message) {
	cmd.originalMessage = message
}

func (cmd *EchoCommand) GetResponseMessage() (*container.Response, error) {
	config := tgbotapi.NewMessage(cmd.originalMessage.Chat.ID, cmd.originalMessage.Text)
	config.ReplyToMessageID = cmd.originalMessage.MessageID
	config.Text = "You just said this."

	response := &container.Response{
		ResponseConfig: config,
		ConfigType: "MessageConfig",
		NoOp: false,
	}

	cmd.hasCompleted = true

	return response, nil
}
