// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCexSellOrder = "cex_sell_order"

// CexSellOrder mapped from table <cex_sell_order>
type CexSellOrder struct {
	ID         int64          `gorm:"column:id;primaryKey" json:"id"`
	Symbol     string         `gorm:"column:symbol;not null" json:"symbol"`
	Price      float64        `gorm:"column:price;not null" json:"price"`
	Quantity   float64        `gorm:"column:quantity;not null" json:"quantity"`
	Status     int32          `gorm:"column:status;not null" json:"status"`
	BuyOrderID int64          `gorm:"column:buy_order_id;not null" json:"buy_order_id"`
	Sold       float64        `gorm:"column:sold;not null" json:"sold"`
	Profit     float64        `gorm:"column:profit;not null" json:"profit"`
	CreateBy   int64          `gorm:"column:create_by;not null" json:"create_by"`
	UpdateBy   int64          `gorm:"column:update_by;not null" json:"update_by"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;not null" json:"deleted_at"`
}

// TableName CexSellOrder's table name
func (*CexSellOrder) TableName() string {
	return TableNameCexSellOrder
}
