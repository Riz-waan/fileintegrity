package main

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.db = initDB()
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 30s", func() { a.UpdateStatus() })

	c.Start()
}

// Add File Dialog
func (a *App) AddFile() {
	value, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return
	}
	hash, err := HashFile(value)
	if err != nil {
		return
	}
	file := Item{LastUpdate: time.Now(), SuccessfulUpdateAt: time.Now(), Location: value, Hash: hash}
	a.db.Create(&file)
	runtime.EventsEmit(a.ctx, "refetch")
}

// Remove File Dialog
func (a *App) RemoveFile(id int) {
	a.db.Delete(&Item{}, id)
	runtime.EventsEmit(a.ctx, "refetch")
}

// Get All Files
func (a *App) GetFiles() []Item {
	var items []Item
	a.db.Find(&items)
	return items
}

func (a *App) UpdateStatus() {
	var items []Item
	a.db.Find(&items, "last_update = successful_update_at")
	for _, item := range items {
		hash, _ := HashFile(item.Location)
		item.LastUpdate = time.Now()
		if hash == item.Hash {
			item.SuccessfulUpdateAt = item.LastUpdate
		}
		a.db.Save(&item)
	}
	runtime.EventsEmit(a.ctx, "refetch")
}
