package main

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Item{})
	return db
}

type Item struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	CreatedAt          time.Time      `json:"-"`
	Location           string         `json:"path"`
	LastUpdate         time.Time      `gorm:"index" json:"updatedAt"`
	SuccessfulUpdateAt time.Time      `gorm:"index" json:"successfulUpdateAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Hash               string         `json:"-"`
}
