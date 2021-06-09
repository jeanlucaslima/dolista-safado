package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object that can be found in an update.
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

func GetMessageHandler(telegramToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody Update
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			logrus.Error("error parsing request body ", err)
			return // TODO: return friendly bot message for error cases
		}

		var telegramApi string = "https://api.telegram.org/bot" + telegramToken + "/sendMessage"
		response, err := http.PostForm(
			telegramApi,
			url.Values{
				"chat_id": {strconv.Itoa(reqBody.Message.Chat.Id)},
				"text":    {"Hello world!"},
			})

		if err != nil {
			logrus.Error("error when posting text to the chat ", err)
			return
		}

		defer response.Body.Close()

		var respBodyBytes, errRead = ioutil.ReadAll(response.Body)
		if errRead != nil {
			logrus.Error("error in parsing telegram answer ", err)
			return
		}

		bodyString := string(respBodyBytes)

		logrus.Info("managed to send message to chat ", reqBody.Message.Chat.Id, bodyString)
	}
}
