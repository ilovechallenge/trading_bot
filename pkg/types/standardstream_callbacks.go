// Code generated by "callbackgen -type StandardStream -interface"; DO NOT EDIT.

package types

import ()

func (s *StandardStream) OnStart(cb func()) {
	s.startCallbacks = append(s.startCallbacks, cb)
}

func (s *StandardStream) EmitStart() {
	for _, cb := range s.startCallbacks {
		cb()
	}
}

func (s *StandardStream) OnConnect(cb func()) {
	s.connectCallbacks = append(s.connectCallbacks, cb)
}

func (s *StandardStream) EmitConnect() {
	for _, cb := range s.connectCallbacks {
		cb()
	}
}

func (s *StandardStream) OnDisconnect(cb func()) {
	s.disconnectCallbacks = append(s.disconnectCallbacks, cb)
}

func (s *StandardStream) EmitDisconnect() {
	for _, cb := range s.disconnectCallbacks {
		cb()
	}
}

func (s *StandardStream) OnAuth(cb func()) {
	s.authCallbacks = append(s.authCallbacks, cb)
}

func (s *StandardStream) EmitAuth() {
	for _, cb := range s.authCallbacks {
		cb()
	}
}

func (s *StandardStream) OnRawMessage(cb func(raw []byte)) {
	s.rawMessageCallbacks = append(s.rawMessageCallbacks, cb)
}

func (s *StandardStream) EmitRawMessage(raw []byte) {
	for _, cb := range s.rawMessageCallbacks {
		cb(raw)
	}
}

func (s *StandardStream) OnTradeUpdate(cb func(trade Trade)) {
	s.tradeUpdateCallbacks = append(s.tradeUpdateCallbacks, cb)
}

func (s *StandardStream) EmitTradeUpdate(trade Trade) {
	for _, cb := range s.tradeUpdateCallbacks {
		cb(trade)
	}
}

func (s *StandardStream) OnOrderUpdate(cb func(order Order)) {
	s.orderUpdateCallbacks = append(s.orderUpdateCallbacks, cb)
}

func (s *StandardStream) EmitOrderUpdate(order Order) {
	for _, cb := range s.orderUpdateCallbacks {
		cb(order)
	}
}

func (s *StandardStream) OnBalanceSnapshot(cb func(balances BalanceMap)) {
	s.balanceSnapshotCallbacks = append(s.balanceSnapshotCallbacks, cb)
}

func (s *StandardStream) EmitBalanceSnapshot(balances BalanceMap) {
	for _, cb := range s.balanceSnapshotCallbacks {
		cb(balances)
	}
}

func (s *StandardStream) OnBalanceUpdate(cb func(balances BalanceMap)) {
	s.balanceUpdateCallbacks = append(s.balanceUpdateCallbacks, cb)
}

func (s *StandardStream) EmitBalanceUpdate(balances BalanceMap) {
	for _, cb := range s.balanceUpdateCallbacks {
		cb(balances)
	}
}

func (s *StandardStream) OnKLineClosed(cb func(kline KLine)) {
	s.kLineClosedCallbacks = append(s.kLineClosedCallbacks, cb)
}

func (s *StandardStream) EmitKLineClosed(kline KLine) {
	for _, cb := range s.kLineClosedCallbacks {
		cb(kline)
	}
}

func (s *StandardStream) OnKLine(cb func(kline KLine)) {
	s.kLineCallbacks = append(s.kLineCallbacks, cb)
}

func (s *StandardStream) EmitKLine(kline KLine) {
	for _, cb := range s.kLineCallbacks {
		cb(kline)
	}
}

func (s *StandardStream) OnBookUpdate(cb func(book SliceOrderBook)) {
	s.bookUpdateCallbacks = append(s.bookUpdateCallbacks, cb)
}

func (s *StandardStream) EmitBookUpdate(book SliceOrderBook) {
	for _, cb := range s.bookUpdateCallbacks {
		cb(book)
	}
}

func (s *StandardStream) OnBookTickerUpdate(cb func(bookTicker BookTicker)) {
	s.bookTickerUpdateCallbacks = append(s.bookTickerUpdateCallbacks, cb)
}

func (s *StandardStream) EmitBookTickerUpdate(bookTicker BookTicker) {
	for _, cb := range s.bookTickerUpdateCallbacks {
		cb(bookTicker)
	}
}

func (s *StandardStream) OnBookSnapshot(cb func(book SliceOrderBook)) {
	s.bookSnapshotCallbacks = append(s.bookSnapshotCallbacks, cb)
}

func (s *StandardStream) EmitBookSnapshot(book SliceOrderBook) {
	for _, cb := range s.bookSnapshotCallbacks {
		cb(book)
	}
}

func (s *StandardStream) OnMarketTrade(cb func(trade Trade)) {
	s.marketTradeCallbacks = append(s.marketTradeCallbacks, cb)
}

func (s *StandardStream) EmitMarketTrade(trade Trade) {
	for _, cb := range s.marketTradeCallbacks {
		cb(trade)
	}
}

func (s *StandardStream) OnAggTrade(cb func(trade Trade)) {
	s.aggTradeCallbacks = append(s.aggTradeCallbacks, cb)
}

func (s *StandardStream) EmitAggTrade(trade Trade) {
	for _, cb := range s.aggTradeCallbacks {
		cb(trade)
	}
}

func (s *StandardStream) OnFuturesPositionUpdate(cb func(futuresPositions FuturesPositionMap)) {
	s.FuturesPositionUpdateCallbacks = append(s.FuturesPositionUpdateCallbacks, cb)
}

func (s *StandardStream) EmitFuturesPositionUpdate(futuresPositions FuturesPositionMap) {
	for _, cb := range s.FuturesPositionUpdateCallbacks {
		cb(futuresPositions)
	}
}

func (s *StandardStream) OnFuturesPositionSnapshot(cb func(futuresPositions FuturesPositionMap)) {
	s.FuturesPositionSnapshotCallbacks = append(s.FuturesPositionSnapshotCallbacks, cb)
}

func (s *StandardStream) EmitFuturesPositionSnapshot(futuresPositions FuturesPositionMap) {
	for _, cb := range s.FuturesPositionSnapshotCallbacks {
		cb(futuresPositions)
	}
}

type StandardStreamEventHub interface {
	OnStart(cb func())

	OnConnect(cb func())

	OnDisconnect(cb func())

	OnAuth(cb func())

	OnRawMessage(cb func(raw []byte))

	OnTradeUpdate(cb func(trade Trade))

	OnOrderUpdate(cb func(order Order))

	OnBalanceSnapshot(cb func(balances BalanceMap))

	OnBalanceUpdate(cb func(balances BalanceMap))

	OnKLineClosed(cb func(kline KLine))

	OnKLine(cb func(kline KLine))

	OnBookUpdate(cb func(book SliceOrderBook))

	OnBookTickerUpdate(cb func(bookTicker BookTicker))

	OnBookSnapshot(cb func(book SliceOrderBook))

	OnMarketTrade(cb func(trade Trade))

	OnAggTrade(cb func(trade Trade))

	OnFuturesPositionUpdate(cb func(futuresPositions FuturesPositionMap))

	OnFuturesPositionSnapshot(cb func(futuresPositions FuturesPositionMap))
}
