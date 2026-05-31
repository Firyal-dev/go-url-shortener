package pkg

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Storage struct {
	db  *gorm.DB
	rdb *redis.Client
}

type URL struct {
	ID      uint   `gorm:"primaryKey"`
	Code    string `gorm:"uniqueIndex;size:6;not null"`
	LongURL string `gorm:"not null"`
}
