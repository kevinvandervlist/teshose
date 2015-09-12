package plugin
import (
	"github.com/kevinvandervlist/teshose/plugin/commands"
)

func (p *Plugin) CreateNoOp() (PluginInstance) {
	return commands.CreateNoOpCommand(p.logger)
}

func (p *Plugin) CreateNoopDescription() string {
	return "noop - Doesn't do anything at all."
}