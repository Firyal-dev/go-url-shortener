package pkg

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"fmt"
)

var ctx = context.Background()

func Connect(db *gorm.DB, rdb *redis.Client) *Storage {
	db.AutoMigrate(&URL{})
	return &Storage{db: db, rdb: rdb}
}

func(s *Storage) Shorten(url string) (string, error) {
	if url == "" {
		return "", errors.New("Url cannot be empty")
	}

	code, _ := GenerateCode(6)
	mapping := URL{
		Code: code,
		LongURL: url,
	}
	result := s.db.Create(&mapping)
	if result.Error != nil {
		return "", result.Error
	}

	s.rdb.Set(ctx, code, url, 24*time.Hour) // 24 hours

	return code, nil
}

func(s *Storage) Resolve(code string) (string, error) {
	cacheURL, err := s.rdb.Get(ctx, code).Result()
	if err == nil {
		return cacheURL, nil
	}
	if err != redis.Nil {
		fmt.Println("Redis error:", err)
	}

	var mapping URL
	result := s.db.Where("code = ?", code).First(&mapping)
	if result.Error != nil {
		return "", result.Error
	}
	s.rdb.Set(ctx, code, mapping.LongURL, 24*time.Hour)

	return mapping.LongURL, nil
}

func GenerateCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}