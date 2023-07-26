package models

import (
	"time"
)

type History struct {
	ID         int                 `gorm:"primaryKey;autoIncrement"`
	UserID     string              `gorm:"index;not null"`
	AnimeID    string              `gorm:"index;not null"`
	AnimeImg   string              `gorm:"not null"`
	AnimeTitle string              `gorm:"not null"`
	CreatedAt  time.Time           `gorm:"default:current_timestamp(6);precision:6"`
	UpdatedAt  time.Time           `gorm:"autoUpdateTime"`
	Episodes   []HistoryPerEpisode `gorm:"foreignKey:HistoryID"`
}

type HistoryPerEpisode struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	EpisodeID string    `gorm:"not null"`
	TimeStamp int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"default:current_timestamp(6);precision:6"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	HistoryID int       `gorm:"index"`
}
