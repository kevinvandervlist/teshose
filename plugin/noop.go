package plugin
import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/container"
)

func (plugin *Plugin) ExecNoOp(incoming *tgbotapi.Message) (*container.Response, error) {
	return &container.Response{
		NoOp: true,
	}, nil
}
