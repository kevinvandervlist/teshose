package plugin
import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/container"
)

func (plugin *Plugin) ExecPhoto(incoming *tgbotapi.Message) (*container.Response, error) {
	config := tgbotapi.NewPhotoUpload(incoming.Chat.ID, "/tmp/image.jpg")
	config.ReplyToMessageID = incoming.MessageID

	response := &container.Response{
		ResponseConfig: config,
		ConfigType: "PhotoConfig",
		NoOp: false,
	}

	return response, nil
}
