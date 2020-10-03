// Code generated by "callbackgen -type StandardStream -interface"; DO NOT EDIT.

package types

import ()

func (stream *StandardStream) OnTrade(cb func(trade *Trade)) {
	stream.tradeCallbacks = append(stream.tradeCallbacks, cb)
}

func (stream *StandardStream) EmitTrade(trade *Trade) {
	for _, cb := range stream.tradeCallbacks {
		cb(trade)
	}
}

func (stream *StandardStream) OnBalanceSnapshot(cb func(balances map[string]Balance)) {
	stream.balanceSnapshotCallbacks = append(stream.balanceSnapshotCallbacks, cb)
}

func (stream *StandardStream) EmitBalanceSnapshot(balances map[string]Balance) {
	for _, cb := range stream.balanceSnapshotCallbacks {
		cb(balances)
	}
}

func (stream *StandardStream) OnBalanceUpdate(cb func(balances map[string]Balance)) {
	stream.balanceUpdateCallbacks = append(stream.balanceUpdateCallbacks, cb)
}

func (stream *StandardStream) EmitBalanceUpdate(balances map[string]Balance) {
	for _, cb := range stream.balanceUpdateCallbacks {
		cb(balances)
	}
}

func (stream *StandardStream) OnKLineClosed(cb func(kline KLine)) {
	stream.kLineClosedCallbacks = append(stream.kLineClosedCallbacks, cb)
}

func (stream *StandardStream) EmitKLineClosed(kline KLine) {
	for _, cb := range stream.kLineClosedCallbacks {
		cb(kline)
	}
}

func (stream *StandardStream) OnBookUpdate(cb func(book OrderBook)) {
	stream.bookUpdateCallbacks = append(stream.bookUpdateCallbacks, cb)
}

func (stream *StandardStream) EmitBookUpdate(book OrderBook) {
	for _, cb := range stream.bookUpdateCallbacks {
		cb(book)
	}
}

func (stream *StandardStream) OnBookSnapshot(cb func(book OrderBook)) {
	stream.bookSnapshotCallbacks = append(stream.bookSnapshotCallbacks, cb)
}

func (stream *StandardStream) EmitBookSnapshot(book OrderBook) {
	for _, cb := range stream.bookSnapshotCallbacks {
		cb(book)
	}
}

type StandardStreamEventHub interface {
	OnTrade(cb func(trade *Trade))

	OnBalanceSnapshot(cb func(balances map[string]Balance))

	OnBalanceUpdate(cb func(balances map[string]Balance))

	OnKLineClosed(cb func(kline KLine))

	OnBookUpdate(cb func(book OrderBook))

	OnBookSnapshot(cb func(book OrderBook))
}
