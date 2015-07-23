package plugin

func (plugin *Plugin) ExecBar(incoming IncomingMessage) (ResponseMessage, error) {
	return ResponseMessage{
		Text: "Bar",
	}, nil
}
