package commands

import (
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/container"
	"github.com/Syfaro/telegram-bot-api"
	"errors"
)

type NoOpCommand struct {
	hasCompleted bool
}

func CreateNoOpCommand(logger *logging.Logger) *NoOpCommand {
	return &NoOpCommand{}
}

func (cmd *NoOpCommand) HasCompleted() bool {
	return true
}

func (cmd *NoOpCommand) SetRequestMessage(message *tgbotapi.Message) {
}

func (cmd *NoOpCommand) GetResponseMessage() (*container.Response, error) {
	return nil, errors.New("NoOp does not do anything.")
}
