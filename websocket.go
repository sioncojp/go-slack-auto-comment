package main

import (
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

// handleMesageEvent ... 受け取ったメッセージのhandle
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) {
	// ループするため、bot自身のメッセージには応答しない
	if ev.Msg.BotID == botID {
		return
	}

	// 指定されたchannelかつ、正規表現にマッチしたワードに対して、response
	for _, v := range s.Actions {
		if ev.Channel == v.ChannelID {
			re := regexp.MustCompile(v.In)
			if re.MatchString(ev.Msg.Text) {
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
