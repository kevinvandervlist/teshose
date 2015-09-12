package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateEcho() (PluginInstance) {
	return commands.CreateEchoCommand(p.logger)
}

func (p *Plugin) CreateEchoDescription() string {
	return "echo - I will repeat everything you say to me."
}