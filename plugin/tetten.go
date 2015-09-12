package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateTetten() (PluginInstance) {
	return commands.CreateTumblrCommand("tettenvrouw", p.logger)
}

func (p *Plugin) CreateTettenDescription() string {
	return "tetten - Does not need an explanation. /tetten 3 for three images."
}