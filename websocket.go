package main

import (
	"fmt"
	"regexp"

	"github.com/nlopes/slack"
)

// SlackListener ... SlackのListenで使うデータを格納
type SlackListener struct {
	Client  *slack.Client
	Actions []Action
}

// NewSlack ...Websocketで受け付けるための、SlackListenerの初期化
func (c *Config) NewSlack() *SlackListener {
	client := slack.New(c.BotToken)
	return &SlackListener{
		Client:  client,
		Actions: c.Actions,
	}
}

// ListenAndResponse ...Websocketの立ち上げ
func (s *SlackListener) ListenAndResponse() {
	rtm := s.Client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			s.handleMessageEvent(ev)
		}
	}
}

// handleMesageEvent ... 受け取ったメッセージのhandle.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) {
	// ループするため、bot自身のメッセージには応答しない
	if ev.Msg.BotID == botID {
		return
	}

	// Textを入れる
	input := ev.Msg.Text

	// Attachmentの場合、title textを入れる
	if ev.Msg.Attachments != nil {
		input = fmt.Sprintf("%s %s", ev.Msg.Attachments[0].Title, ev.Msg.Attachments[0].Text)

		if ev.Msg.Attachments[0].Fields != nil {
			for _, v := range ev.Msg.Attachments[0].Fields {
				input = fmt.Sprintf("%s\n%s\n%s", input, v.Title, v.Value)
			}
		}
	}

	// 指定されたchannelかつ、正規表現にマッチしたワードに対して、response
	for _, v := range s.Actions {
		if v.ChannelID == "" || ev.Channel == v.ChannelID {
			re, err := regexp.Compile(v.In)
			if err != nil {
				continue
			}
			if re.MatchString(input) {
				s.ResponseMessage(ev.Msg.Timestamp, ev.Channel, v.Out)
			}
		}
	}
}

// ResponseMessage ... response to target thread in Slack
func (s *SlackListener) ResponseMessage(ts, channel, msg string) {
	if _, _, err := s.Client.PostMessage(channel, slack.MsgOptionTS(ts), slack.MsgOptionText(msg, false)); err != nil {
		log.sugar.Errorf("failed to post message: %s", err)
	}
}
