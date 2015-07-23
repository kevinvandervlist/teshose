package plugin

import (
	"reflect"
	"errors"
	"strings"
	"github.com/op/go-logging"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/container"
)

type Plugin struct {
	logger *logging.Logger
}

func Create(logger *logging.Logger) (*Plugin) {
	return &Plugin{
		logger: logger,
	}
}

// Invoke this method with the name of the actual method you want to invoke. The second parameter should be the
// received message. The message to send as a response should be returned by it.
func (plugin *Plugin) Exec(cmd string, message *tgbotapi.Message) (*container.Response, error) {
	pluginCmd, err := plugin.getCommand(cmd)
	plugin.logger.Debug("Exctracted command: %s", cmd)

	if(err != nil) {
		plugin.logger.Debug("No command found")
		return plugin.ExecNoOp(message)
	}

	_func, err := plugin.getMethod(pluginCmd)

	if(err != nil) {
		plugin.logger.Debug("Command %s has no plugin -- noop.", pluginCmd)
		return plugin.ExecNoOp(message)
	} else {
		plugin.logger.Debug("Executing command %s.", pluginCmd)
		return _func(message)

	}
}

func (plugin *Plugin) getMethod(pluginCmd string) (func(*tgbotapi.Message) (*container.Response, error), error) {
	method := reflect.ValueOf(plugin).MethodByName(pluginCmd)
	if !method.IsValid() {
		return nil, errors.New("No plugin found")
	}
	return method.Interface().(func(*tgbotapi.Message)(*container.Response, error)), nil
}

func (plugin *Plugin) getCommand(cmd string) (string, error) {
	prefix := "Exec"
	plugin.logger.Debug("Input string: '%s'", cmd)
	e := "Not a valid input string"

	splitted := strings.Split(cmd, " ")
	if len(splitted) == 0 {
		plugin.logger.Debug("Invalid input!")
		return "", errors.New(e)
	}

	if (strings.HasPrefix(splitted[0], "/") || strings.HasPrefix(splitted[0], "@")) {
		t := splitted[0][1:]
		before := prefix + strings.ToUpper(t[:1]) + strings.ToLower(t[1:])
		return before, nil
	} else {
		return "", errors.New(e)
	}
}