// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameChainbotBotTemplate = "chainbot_bot_template"

// ChainbotBotTemplate mapped from table <chainbot_bot_template>
type ChainbotBotTemplate struct {
	ID             int64          `gorm:"column:id;primaryKey" json:"id"`
	DriverType     string         `gorm:"column:driver_type;not null" json:"driver_type"`
	NotifyTemplate string         `gorm:"column:notify_template;not null" json:"notify_template"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;not null" json:"deleted_at"`
}

// TableName ChainbotBotTemplate's table name
func (*ChainbotBotTemplate) TableName() string {
	return TableNameChainbotBotTemplate
}
