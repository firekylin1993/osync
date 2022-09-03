package model

import "time"

type OsyncModel struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	AppName       string    `json:"app_name"`
	ChannelId     int       `json:"channel_id"`
	ChannelName   string    `json:"channel_name"`
	PackageId     int       `json:"package_id"`
	UpdateVersion string    `json:"update_version"`
	Status        int       `gorm:"default:0" json:"status"`
	CreateDate    time.Time `gorm:"autoCreateTime" json:"create_date"`
	UpdateDate    time.Time `gorm:"autoUpdateTime" json:"update_date"`
}
