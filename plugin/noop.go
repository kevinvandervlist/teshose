package plugin
import "github.com/kevinvandervlist/teshose/messages"

func (plugin *Plugin) ExecNoOp(incoming *messages.IncomingMessage) (*messages.ResponseMessage, error) {
	return &messages.ResponseMessage{
		NoOp: true,
	}, nil
}
