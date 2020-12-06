package telegramnotifier

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Notifier struct {
	Bot      *tb.Bot
	chatUser *tb.User
	channel  string
}

type NotifyOption func(notifier *Notifier)

// start bot daemon
func New(botToken, initToken string, options ...NotifyOption) *Notifier {
	bot, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		// URL: "http://195.129.111.17:8012",

		Token:  botToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}

	chatUser := &tb.User{}

	bot.Handle("/help", func(m *tb.Message) {
		helpMsg := `
help	- print help message
init	- initialize telegram bot with initToken. ex. /init my-token
hi		- print welcome message to authorized user
`
		bot.Send(m.Sender, helpMsg)
	})

	// init check initToken and then set sender id
	bot.Handle("/init", func(m *tb.Message) {
		log.Info("Receive message: ", m) //debug
		if m.Payload == initToken {
			chatUser = m.Sender
			bot.Send(m.Sender, "Bot initialized")
		} else {
			bot.Send(m.Sender, "Error: bot intialize failed. Init token not match!")
		}
	})

	bot.Handle("/hi", func(m *tb.Message) {
		if m.Sender.ID == chatUser.ID {
			bot.Send(chatUser,
				fmt.Sprintf("Welcome! user: %s, ID: %s"),
				chatUser.Username,
				chatUser.ID)
		} else {
			log.Warningf("Incorrect user tried to access bot! sender id: %s", m.Sender)
		}
	})

	notifier := &Notifier{
		chatUser: chatUser,
		Bot:      bot,
	}

	for _, o := range options {
		o(notifier)
	}

	return notifier
}

func (n *Notifier) Notify(format string, args ...interface{}) {
	n.NotifyTo(n.channel, format, args...)
}

func (n *Notifier) NotifyTo(channel, format string, args ...interface{}) {
	var telegramArgsOffset = -1

	var nonTelegramArgs = args
	if telegramArgsOffset > -1 {
		nonTelegramArgs = args[:telegramArgsOffset]
	}

	log.Infof(format, nonTelegramArgs...)

	_, err := n.Bot.Send(n.chatUser, fmt.Sprintf(format, nonTelegramArgs...))
	if err != nil {
		log.WithError(err).
			WithField("chatUser", n.chatUser).
			Errorf("telegram error: %s", err.Error())
	}

	return
}
