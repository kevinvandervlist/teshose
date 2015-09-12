package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateLingerie() (PluginInstance) {
	return commands.CreateTumblrCommand("lingeriebomb", p.logger)
}

func (p *Plugin) CreateLingerieDescription() string {
	return "lingerie - Send some lingerie photos. /lingerie 3 will give you three of those images."
}