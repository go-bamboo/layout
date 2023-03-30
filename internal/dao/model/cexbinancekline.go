package model

type CexBinanceKline struct {
	Event                string  `json:"event"`
	StartTime            int     `json:"start_time"`
	EndTime              int     `json:"end_time"`
	Symbol               string  `json:"symbol"`
	Interval             string  `json:"interval"`
	FirstTradeID         int     `json:"first_trade_id"`
	LastTradeID          int     `json:"last_trade_id"`
	Open                 float64 `json:"open"`
	Close                float64 `json:"close"`
	High                 float64 `json:"high"`
	Low                  float64 `json:"low"`
	Volume               string  `json:"volume"`
	TradeNum             int     `json:"trade_num"`
	IsFinal              bool    `json:"is_final"`
	QuoteVolume          string  `json:"quote_volume"`
	ActiveBuyVolume      string  `json:"active_buy_volume"`
	ActiveBuyQuoteVolume string  `json:"active_buy_quote_volume"`
}

func (CexBinanceKline) TableName() string {
	return "cex_binance_kline"
}
