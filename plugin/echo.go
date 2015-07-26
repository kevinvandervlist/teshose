package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateEcho() (PluginInstance) {
	return commands.CreateEchoCommand(p.logger)
}
