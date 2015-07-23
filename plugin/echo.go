package plugin

func (plugin *Plugin) ExecEcho(incoming IncomingMessage) (ResponseMessage, error) {
	return ResponseMessage{
		Text: incoming.Text,
	}, nil
}
