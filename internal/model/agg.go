package model

import "time"

type ChannelPackageModel struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	ChannelName   string    `json:"channel_name"`
	UpdateVersion string    `json:"update_version"`
	CreateDate    time.Time `gorm:"autoCreateTime" json:"create_date"`
	UpdateDate    time.Time `gorm:"autoUpdateTime" json:"update_date"`
}
