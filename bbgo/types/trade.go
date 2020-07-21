package types

import (
	"github.com/slack-go/slack"
	"time"
)

type Trade struct {
	// GID is the global ID
	GID int64 `json:"gid" db:"gid"`

	// ID is the source trade ID
	ID          int64     `json:"id" db:"id"`
	Exchange    string    `json:"exchange" db:"exchange"`
	Price       float64   `json:"price" db:"price"`
	Volume      float64   `json:"volume" db:"volume"`
	Symbol      string    `json:"symbol" db:"symbol"`
	Side        string    `json:"side" db:"side"`
	IsBuyer     bool      `json:"isBuyer" db:"is_buyer"`
	IsMaker     bool      `json:"isMaker" db:"is_maker"`
	Time        time.Time `json:"tradedAt" db:"traded_at"`
	Fee         float64   `json:"fee" db:"fee"`
	FeeCurrency string    `json:"feeCurrency" db:"fee_currency"`
}

func (trade Trade) SlackAttachment() slack.Attachment {
	var color = ""
	if trade.IsBuyer {
		color = "#228B22"
	} else {
		color = "#DC143C"
	}

	market, ok := FindMarket(trade.Symbol)
	if !ok {
		return slack.Attachment{
			Title: "New Trade",
			Color: color,
		}
	}

	return slack.Attachment{
		Title: "New Trade",
		Color: color,
		// Pretext:       "",
		// Text:          "",
		Fields: []slack.AttachmentField{
			{Title: "Symbol", Value: trade.Symbol, Short: true},
			{Title: "Side", Value: trade.Side, Short: true},
			{Title: "Price", Value: market.FormatPrice(trade.Price), Short: true},
			{Title: "Volume", Value: market.FormatVolume(trade.Volume), Short: true},
		},
		// Footer:     tradingCtx.TradeStartTime.Format(time.RFC822),
		// FooterIcon: "",
	}
}
