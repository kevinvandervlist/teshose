package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateHelp() (PluginInstance) {
	return commands.CreateHelpCommand(p.logger)
}

func (p *Plugin) CreateHelpDescription() string {
	return "help - Usage information."
}