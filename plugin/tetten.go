package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateTetten() (PluginInstance) {
	return commands.CreateTettenVrouwCommand(p.logger)
}