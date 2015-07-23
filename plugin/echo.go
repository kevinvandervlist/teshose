package plugin
import "github.com/kevinvandervlist/teshose/messages"

func (plugin *Plugin) ExecEcho(incoming *messages.IncomingMessage) (*messages.ResponseMessage, error) {
	response := messages.Convert(incoming)
	response.Text = "I'll just echo your messages."
	return response, nil
}
