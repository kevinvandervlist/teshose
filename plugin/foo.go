package plugin

func (plugin *Plugin) ExecFoo(incoming IncomingMessage) (ResponseMessage, error) {
	return ResponseMessage{
		Text: "Foo",
	}, nil
}
