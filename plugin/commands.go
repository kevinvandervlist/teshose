package plugin

import (
	"reflect"
	"errors"
	"strings"
	"github.com/op/go-logging"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/container"
	"sort"
	"fmt"
)

type Plugin struct {
	logger *logging.Logger
}

type PluginInstance interface {
	HasCompleted() bool
	SetRequestMessage(message *tgbotapi.Message)
	GetResponseMessage() (*container.Response, error)
}

func Create(logger *logging.Logger) (*Plugin) {
	return &Plugin{
		logger: logger,
	}
}

func (plugin *Plugin) CreateDescriptions() {
	cmds := []string{"echo", "help", "lingerie", "tetten", "noop"}
	sort.Strings(cmds)

	plugin.logger.Info("Available descriptions: ")

	for _,cmd := range cmds {
		_func, err := plugin.getDescription(cmd)
		if (err != nil) {
			panic(err)
		}
		fmt.Printf("%s\n", _func())
	}
}

func (plugin *Plugin) BuildPluginInstance(cmd string) PluginInstance {
	pluginCmd, err := plugin.getCommand(cmd)
	plugin.logger.Debug("Exctracted command: %s", cmd)

	if(err != nil) {
		plugin.logger.Debug("No command found")
		return plugin.CreateNoOp()
	}

	_func, err := plugin.getMethod(pluginCmd)

	if(err != nil) {
		plugin.logger.Debug("Command %s has no plugin -- noop.", pluginCmd)
		return plugin.CreateNoOp()
	} else {
		plugin.logger.Debug("Executing command %s.", pluginCmd)
		return _func()

	}
}

func (plugin *Plugin) getMethod(pluginCmd string) (func() (PluginInstance), error) {
	method := reflect.ValueOf(plugin).MethodByName(pluginCmd)
	if !method.IsValid() {
		return nil, errors.New("No plugin found")
	}
	return method.Interface().(func()(PluginInstance)), nil
}

func (plugin *Plugin) getDescription(cmd string) (func() (string), error) {
	before := "Create" + strings.ToUpper(cmd[:1]) + strings.ToLower(cmd[1:]) + "Description"
	method := reflect.ValueOf(plugin).MethodByName(before)
	if !method.IsValid() {
		return nil, errors.New("No plugin found")
	}
	return method.Interface().(func()(string)), nil
}

func (plugin *Plugin) getCommand(cmd string) (string, error) {
	prefix := "Create"
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