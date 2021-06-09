package main

import (
	"net/http"
	"os"
	"time"

	"github.com/iatistas/dolista-safado/service"
	"github.com/sirupsen/logrus"
)

const noToken = "no-token"

func main() {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	if telegramToken == "" || telegramToken == noToken {
		logrus.Error("telegram token not provided")
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("/message", service.GetMessageHandler(telegramToken))

	server := &http.Server{
		Addr:         ":80",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	logrus.Info("starting server on port 80")
	logrus.Error(server.ListenAndServe())
}
