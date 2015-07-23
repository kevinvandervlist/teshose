package api

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/container"
)

type TelegramApi struct {
	token string
	debug bool
	currentUpdateID int
	longPollingTimeout int
	bot *tgbotapi.BotAPI
	disconnect chan bool
	ReceiveMessagesChannel chan *tgbotapi.Message
	SendMessagesChannel chan *container.Response
	ErrorChannel chan error
	logger *logging.Logger
}

func Create(logger *logging.Logger, token string) (*TelegramApi) {
	return &TelegramApi{
		token,
		false,
		0,
		30,
		nil,
		make(chan bool, 1),
		make(chan *tgbotapi.Message, 100),
		make(chan *container.Response, 100),
		make(chan error, 10),
		logger,
	}
}

func (api *TelegramApi) Connect() (error) {
	bot, err := tgbotapi.NewBotAPI(api.token)
	if(err != nil) {
		return err
	}
	api.bot = bot

	go func() {
		for {
			select {
			case <- api.disconnect:
				api.logger.Debug("Disconnected fetcher!")
				return
			default:
				config := tgbotapi.NewUpdate(api.currentUpdateID + 1)
				config.Timeout = api.longPollingTimeout

				updates, err := api.bot.GetUpdates(config)
				api.logger.Debug("GetUpdates() done, received %d messages", len(updates))

				if(err != nil) {
					api.ErrorChannel <- err
				} else {
					for _, update := range updates {
						api.currentUpdateID = update.UpdateID
						var message tgbotapi.Message = update.Message
						api.ReceiveMessagesChannel <- &message
					}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <- api.disconnect:
				api.logger.Debug("Disconnected producer!")
				return
			default:
				response := <- api.SendMessagesChannel
				switch response.ConfigType {
					case "MessageConfig":
						config, ok := response.ResponseConfig.(tgbotapi.MessageConfig)
						if ok {
							_, err := api.bot.SendMessage(config)
							if(err != nil) {
								api.logger.Error("An error occurred while sending a message.", err)
							}
						}
					case "PhotoConfig":
						config, ok := response.ResponseConfig.(tgbotapi.PhotoConfig)
						if ok {
							_, err := api.bot.SendPhoto(config)
							if(err != nil) {
								api.logger.Error("An error occurred while sending a message.", err)
							}
						}
					default:
						api.logger.Error("No implementation found for type %s", response.ConfigType)
				}
			}
		}
	}()

	return nil
}

func (api *TelegramApi) Disconnect() {
	api.logger.Debug("Disconnect requested.")
	api.disconnect <- true
}

func (api *TelegramApi) Debug(state bool) {
	api.debug = state
}

func (api *TelegramApi) GetMe() (tgbotapi.User, error) {
	return api.bot.GetMe()
}

