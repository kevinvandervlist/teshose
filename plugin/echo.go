package plugin
import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/container"
)

func (plugin *Plugin) ExecEcho(incoming *tgbotapi.Message) (*container.Response, error) {
	config := tgbotapi.NewMessage(incoming.Chat.ID, incoming.Text)
	config.ReplyToMessageID = incoming.MessageID
	config.Text = "I'll just echo your messages."

	response := &container.Response{
		ResponseConfig: config,
		ConfigType: "MessageConfig",
		NoOp: false,
	}

	return response, nil
}
