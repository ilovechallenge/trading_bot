package bbgo

import (
	"fmt"
	"sync"
	"time"

	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/types"
	"github.com/c9s/bbgo/pkg/util"
	"github.com/slack-go/slack"
)

type ExchangeFee struct {
	MakerFeeRate fixedpoint.Value
	TakerFeeRate fixedpoint.Value
}

type Position struct {
	Symbol        string `json:"symbol"`
	BaseCurrency  string `json:"baseCurrency"`
	QuoteCurrency string `json:"quoteCurrency"`

	Base        fixedpoint.Value `json:"base"`
	Quote       fixedpoint.Value `json:"quote"`
	AverageCost fixedpoint.Value `json:"averageCost"`

	ExchangeFeeRates map[types.ExchangeName]ExchangeFee `json:"exchangeFeeRates"`

	sync.Mutex
}

func (p *Position) Reset() {
	p.Base = 0
	p.Quote = 0
	p.AverageCost = 0
}

func (p *Position) SetExchangeFeeRate(ex types.ExchangeName, exchangeFee ExchangeFee) {
	if p.ExchangeFeeRates == nil {
		p.ExchangeFeeRates = make(map[types.ExchangeName]ExchangeFee)
	}

	p.ExchangeFeeRates[ex] = exchangeFee
}

func (p *Position) SlackAttachment() slack.Attachment {
	p.Lock()
	averageCost := p.AverageCost
	base := p.Base
	quote := p.Quote
	p.Unlock()

	var posType = ""
	var color = ""

	if p.Base == 0 {
		color = "#cccccc"
		posType = "Closed"
	} else if p.Base > 0 {
		posType = "Long"
		color = "#228B22"
	} else if p.Base < 0 {
		posType = "Short"
		color = "#DC143C"
	}

	title := util.Render(posType+` Position {{ .Symbol }} `, p)
	return slack.Attachment{
		// Pretext:       "",
		// Text:  text,
		Title: title,
		Color: color,
		Fields: []slack.AttachmentField{
			{Title: "Average Cost", Value: util.FormatFloat(averageCost.Float64(), 2), Short: true},
			{Title: p.BaseCurrency, Value: util.FormatFloat(base.Float64(), 4), Short: true},
			{Title: p.QuoteCurrency, Value: util.FormatFloat(quote.Float64(), 2)},
		},
		Footer: util.Render("update time {{ . }}", time.Now().Format(time.RFC822)),
		// FooterIcon: "",
	}
}

func (p *Position) PlainText() string {
	return fmt.Sprintf("Position %s: average cost = %f, base = %f, quote = %f",
		p.Symbol,
		p.AverageCost.Float64(),
		p.Base.Float64(),
		p.Quote.Float64(),
	)
}

func (p *Position) String() string {
	return fmt.Sprintf("POSITION %s: average cost = %f, base = %f, quote = %f",
		p.Symbol,
		p.AverageCost.Float64(),
		p.Base.Float64(),
		p.Quote.Float64(),
	)
}

func (p *Position) BindStream(stream types.Stream) {
	stream.OnTradeUpdate(func(trade types.Trade) {
		if p.Symbol == trade.Symbol {
			p.AddTrade(trade)
		}
	})
}

func (p *Position) AddTrades(trades []types.Trade) (fixedpoint.Value, bool) {
	var totalProfitAmount fixedpoint.Value
	for _, trade := range trades {
		if profitAmount, profit := p.AddTrade(trade); profit {
			totalProfitAmount += profitAmount
		}
	}

	return totalProfitAmount, totalProfitAmount != 0
}

func (p *Position) AddTrade(t types.Trade) (fixedpoint.Value, bool) {
	price := fixedpoint.NewFromFloat(t.Price)
	quantity := fixedpoint.NewFromFloat(t.Quantity)
	quoteQuantity := fixedpoint.NewFromFloat(t.QuoteQuantity)
	fee := fixedpoint.NewFromFloat(t.Fee)

	// calculated fee in quote (some exchange accounts may enable platform currency fee discount, like BNB)
	var quoteFee fixedpoint.Value = 0

	switch t.FeeCurrency {

	case p.BaseCurrency:
		quantity -= fee

	case p.QuoteCurrency:
		quoteQuantity -= fee

	default:
		if p.ExchangeFeeRates != nil {
			if exchangeFee, ok := p.ExchangeFeeRates[t.Exchange]; ok {
				if t.IsMaker {
					quoteFee += exchangeFee.MakerFeeRate.Mul(quoteQuantity)
				} else {
					quoteFee += exchangeFee.TakerFeeRate.Mul(quoteQuantity)
				}
			}
		}
	}

	p.Lock()
	defer p.Unlock()

	// Base > 0 means we're in long position
	// Base < 0  means we're in short position
	switch t.Side {

	case types.SideTypeBuy:
		if p.Base < 0 {
			// handling short-to-long position
			if p.Base+quantity > 0 {
				closingProfit := (p.AverageCost - price).Mul(-p.Base) - quoteFee
				p.Base += quantity
				p.Quote -= quoteQuantity
				p.AverageCost = price
				return closingProfit, true
			} else {
				// covering short position
				p.Base += quantity
				p.Quote -= quoteQuantity
				return (p.AverageCost - price).Mul(quantity) - quoteFee, true
			}
		}

		p.AverageCost = (p.AverageCost.Mul(p.Base) + quoteQuantity + quoteFee).Div(p.Base + quantity)
		p.Base += quantity
		p.Quote -= quoteQuantity

		return 0, false

	case types.SideTypeSell:
		if p.Base > 0 {
			// long-to-short
			if p.Base-quantity < 0 {
				closingProfit := (price - p.AverageCost).Mul(p.Base) - quoteFee
				p.Base -= quantity
				p.Quote += quoteQuantity
				p.AverageCost = price
				return closingProfit, true
			} else {
				p.Base -= quantity
				p.Quote += quoteQuantity
				return (price - p.AverageCost).Mul(quantity) - quoteFee, true
			}
		}

		// handling short position, since Base here is negative we need to reverse the sign
		p.AverageCost = (p.AverageCost.Mul(-p.Base) + quoteQuantity - quoteFee).Div(-p.Base + quantity)
		p.Base -= quantity
		p.Quote += quoteQuantity

		return 0, false
	}

	return 0, false
}
