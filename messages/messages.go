package messages
import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/kevinvandervlist/teshose/api"
)

// TODO: Fix this mess.

type MessageBroker struct {
	api *api.TelegramApi
}

type IncomingMessage struct {
	Text string
	ChatID int
	MessageID int
}

type ResponseMessage struct {
	Text string
	NoOp bool
	inbound *IncomingMessage
}

func Create(api *api.TelegramApi) *MessageBroker {
	return &MessageBroker{
		api: api,
	}
}

func Convert(incoming *IncomingMessage) *ResponseMessage {
	return &ResponseMessage{
		NoOp: false,
		inbound: incoming,
	}
}

func (broker *MessageBroker) ConvertInbound(inbound *tgbotapi.Message) IncomingMessage {
	return IncomingMessage{
		Text: inbound.Text,
		ChatID: inbound.Chat.ID,
		MessageID: inbound.MessageID,
	}
}

func (broker *MessageBroker) ConvertOutbound(outbound *ResponseMessage) *tgbotapi.MessageConfig {
	msg := broker.api.NewMessage(outbound.inbound.ChatID, outbound.Text)
	msg.ReplyToMessageID = outbound.inbound.MessageID
	return &msg
}
