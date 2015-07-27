package commands

import (
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/container"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/plugin/backends"
	"os"
	"strings"
	"strconv"
)

type TettenVrouwCommand struct {
	hasCompleted bool
	logger *logging.Logger
	originalMessage *tgbotapi.Message
	expected int
	handled int
}

func CreateTettenVrouwCommand(logger *logging.Logger) *TettenVrouwCommand {
	return &TettenVrouwCommand{
		logger: logger,
		expected: 1,
		handled: 0,
	}
}

func (cmd *TettenVrouwCommand) HasCompleted() bool {
	return cmd.handled == cmd.expected
}

func (cmd *TettenVrouwCommand) SetRequestMessage(message *tgbotapi.Message) {
	cmd.originalMessage = message
	splitted := strings.Split(message.Text, " ")
	if(len(splitted) < 2) {
		cmd.expected = 1
	} else {
		number, err := strconv.Atoi(splitted[1])
		if err != nil {
			cmd.expected = 1
		} else {
			cmd.expected = number
		}
	}

	if cmd.expected >= 15 {
		cmd.expected = 15
	}
}

func (cmd *TettenVrouwCommand) GetResponseMessage() (*container.Response, error) {
	tumblr := backends.CreateTumblr("tettenvrouw")
	page, err := tumblr.GetRandomPage()
	if err != nil {
		return nil, err
	}

	url, err := tumblr.GetImageUrlFromPage(page)
	if err != nil {
		return nil, err
	}
	cmd.logger.Debug("Download URL %s\n", url)

	path, err := tumblr.DownloadImage(url)
	if err != nil {
		cmd.logger.Debug("Error while writing URL %s\n", url)
		return nil, err
	}
	cmd.logger.Debug("Downloaded image %s to path %s\n", url, path)

	config := tgbotapi.NewPhotoUpload(cmd.originalMessage.Chat.ID, path)
	config.ReplyToMessageID = cmd.originalMessage.MessageID
	config.Caption = tumblr.GetName()

	response := &container.Response{
		ResponseConfig: config,
		ConfigType: "PhotoConfig",
		CallBack: func() {
			os.Remove(path)
		},
	}

	cmd.handled += 1

	return response, nil
}
