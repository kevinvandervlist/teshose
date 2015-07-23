package plugin

import (
	"reflect"
	"errors"
)

type Plugin struct {

}

type IncomingMessage struct {
	Text string
}

type ResponseMessage struct {
	Text string
}

func Create() (*Plugin) {
	return &Plugin{}
}

// Invoke this method with the name of the actual method you want to invoke. The second parameter should be the
// received message. The message to send as a response should be returned by it.
func (plugin *Plugin) Exec(cmd string, message IncomingMessage) (ResponseMessage, error) {
	pluginCmd := "Exec" + cmd
	_func, err := plugin.getMethod(pluginCmd)
	if(err != nil) {
		return ResponseMessage{}, err
	} else {
		return _func(message)

	}

}

func (plugin *Plugin) getMethod(pluginCmd string) (func(IncomingMessage) (ResponseMessage, error), error) {
	method := reflect.ValueOf(plugin).MethodByName(pluginCmd)
	if !method.IsValid() {
		return nil, errors.New("No plugin found")
	}
	return method.Interface().(func(IncomingMessage)(ResponseMessage, error)), nil
}
