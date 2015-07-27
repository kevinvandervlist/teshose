package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateHelp() (PluginInstance) {
	return commands.CreateHelpCommand(p.logger)
}