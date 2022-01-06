package types

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/slack-go/slack"

	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/util"
)

func init() {
	// make sure we can cast Trade to PlainText
	_ = PlainText(Trade{})
	_ = PlainText(&Trade{})
}

type TradeSlice struct {
	mu     sync.Mutex
	Trades []Trade
}

func (s *TradeSlice) Copy() []Trade {
	s.mu.Lock()
	slice := make([]Trade, len(s.Trades), len(s.Trades))
	copy(slice, s.Trades)
	s.mu.Unlock()

	return slice
}

func (s *TradeSlice) Reverse() {
	slice := s.Trades
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func (s *TradeSlice) Append(t Trade) {
	s.mu.Lock()
	s.Trades = append(s.Trades, t)
	s.mu.Unlock()
}

type Trade struct {
	// GID is the global ID
	GID int64 `json:"gid" db:"gid"`

	// ID is the source trade ID
	ID            uint64       `json:"id" db:"id"`
	OrderID       uint64       `json:"orderID" db:"order_id"`
	Exchange      ExchangeName `json:"exchange" db:"exchange"`
	Price         float64      `json:"price" db:"price"`
	Quantity      float64      `json:"quantity" db:"quantity"`
	QuoteQuantity float64      `json:"quoteQuantity" db:"quote_quantity"`
	Symbol        string       `json:"symbol" db:"symbol"`

	Side        SideType `json:"side" db:"side"`
	IsBuyer     bool     `json:"isBuyer" db:"is_buyer"`
	IsMaker     bool     `json:"isMaker" db:"is_maker"`
	Time        Time     `json:"tradedAt" db:"traded_at"`
	Fee         float64  `json:"fee" db:"fee"`
	FeeCurrency string   `json:"feeCurrency" db:"fee_currency"`

	IsMargin   bool `json:"isMargin" db:"is_margin"`
	IsFutures  bool `json:"isFutures" db:"is_futures"`
	IsIsolated bool `json:"isIsolated" db:"is_isolated"`

	StrategyID sql.NullString  `json:"strategyID" db:"strategy"`
	PnL        sql.NullFloat64 `json:"pnl" db:"pnl"`
}

func (trade Trade) PositionChange() fixedpoint.Value {
	q := fixedpoint.NewFromFloat(trade.Quantity)
	switch trade.Side {
	case SideTypeSell:
		return -q

	case SideTypeBuy:
		return q

	case SideTypeSelf:
		return 0
	}
	return 0
}

func trimTrailingZero(a string) string {
	index := strings.Index(a, ".")
	if index == -1 {
		return a
	}

	var c byte
	var i int
	for i = len(a) - 1; i >= 0; i-- {
		c = a[i]
		if c == '0' {
			continue
		} else if c == '.' {
			return a[0:i]
		} else {
			return a[0 : i+1]
		}
	}
	return a
}

func trimTrailingZeroFloat(a float64) string {
	return trimTrailingZero(fmt.Sprintf("%f", a))
}

// String is for console output
func (trade Trade) String() string {
	return fmt.Sprintf("TRADE %s %s %4s %s @ %s amount %s fee %s %s orderID %d %s",
		trade.Exchange.String(),
		trade.Symbol,
		trade.Side,
		trimTrailingZeroFloat(trade.Quantity),
		trimTrailingZeroFloat(trade.Price),
		trimTrailingZeroFloat(trade.QuoteQuantity),
		trimTrailingZeroFloat(trade.Fee),
		trade.FeeCurrency,
		trade.OrderID,
		trade.Time.Time().Format(time.StampMilli),
	)
}

// PlainText is used for telegram-styled messages
func (trade Trade) PlainText() string {
	return fmt.Sprintf("Trade %s %s %s %s @ %s, amount %s, fee %s %s",
		trade.Exchange.String(),
		trade.Symbol,
		trade.Side,
		trimTrailingZeroFloat(trade.Quantity),
		trimTrailingZeroFloat(trade.Price),
		trimTrailingZeroFloat(trade.QuoteQuantity),
		trimTrailingZeroFloat(trade.Fee),
		trade.FeeCurrency)
}

var slackTradeTextTemplate = ":handshake: {{ .Symbol }} {{ .Side }} Trade Execution @ {{ .Price  }}"

func (trade Trade) SlackAttachment() slack.Attachment {
	var color = "#DC143C"

	if trade.IsBuyer {
		color = "#228B22"
	}

	liquidity := trade.Liquidity()
	text := util.Render(slackTradeTextTemplate, trade)
	return slack.Attachment{
		Text: text,
		// Title: ...
		// Pretext: pretext,
		Color: color,
		Fields: []slack.AttachmentField{
			{Title: "Exchange", Value: trade.Exchange.String(), Short: true},
			{Title: "Price", Value: trimTrailingZeroFloat(trade.Price), Short: true},
			{Title: "Quantity", Value: trimTrailingZeroFloat(trade.Quantity), Short: true},
			{Title: "QuoteQuantity", Value: trimTrailingZeroFloat(trade.QuoteQuantity), Short: true},
			{Title: "Fee", Value: trimTrailingZeroFloat(trade.Fee), Short: true},
			{Title: "FeeCurrency", Value: trade.FeeCurrency, Short: true},
			{Title: "Liquidity", Value: liquidity, Short: true},
			{Title: "Order ID", Value: strconv.FormatUint(trade.OrderID, 10), Short: true},
		},
		Footer: util.Render("trade time {{ . }}", trade.Time.Time().Format(time.StampMilli)),
	}
}

func (trade Trade) Liquidity() (o string) {
	if trade.IsMaker {
		o = "MAKER"
	} else {
		o = "TAKER"
	}

	return o
}

func (trade Trade) Key() TradeKey {
	return TradeKey{
		Exchange: trade.Exchange,
		ID:       trade.ID,
		Side:     trade.Side,
	}
}

type TradeKey struct {
	Exchange ExchangeName
	ID       uint64
	Side     SideType
}
