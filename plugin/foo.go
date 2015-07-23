package plugin
import "github.com/kevinvandervlist/teshose/messages"

func (plugin *Plugin) ExecFoo(incoming *messages.IncomingMessage) (*messages.ResponseMessage, error) {
	response := messages.Convert(incoming)
	response.Text = "Appel, Banaan"
	return response, nil
}
