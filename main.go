package main

import (
	"fmt"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	slackClient *slack.Client
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("error reading config %v", err)
	}

	slackAccessKey := viper.GetString("solobot.slackAccessKey")
	log.Println(slackAccessKey)
	slackClient = slack.New(slackAccessKey)
	rtm := slackClient.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch evt := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(evt)
		}
	}
}

func handleMessage(evt *slack.MessageEvent) {
	fmt.Println(evt.Msg.Text)
	replyToUser(evt)
}

func replyToUser(evt *slack.MessageEvent) {
	reply := slack.MsgOptionText("Hi, there", true)
	toUser := slack.MsgOptionAsUser(true)
	slackClient.PostMessage(evt.User, reply, toUser)
}
