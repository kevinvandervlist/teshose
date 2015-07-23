package plugin
import "github.com/kevinvandervlist/teshose/messages"

func (plugin *Plugin) ExecBar(incoming *messages.IncomingMessage) (*messages.ResponseMessage, error) {
	response := messages.Convert(incoming)
	response.Text = "Koelkast, Televisie, Tafel, ..."
	return response, nil
}
