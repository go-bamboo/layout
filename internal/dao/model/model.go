package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type SingResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type TradeStat struct {
	BuyerAmount  float64
	SellerAmount float64
}

type OrderStat struct {
	Price    float64
	Quantity float64
	CreateAt time.Time
}

type ProfitStat struct {
}

type OrderCmd struct {
	Cmd string
}
