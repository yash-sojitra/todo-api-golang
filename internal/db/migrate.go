package db

import (
	"github.com/yash-sojitra/todo/internal/models"
	"gorm.io/gorm"
)

type MigrationHandler struct {
	DB *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{DB: db}
}

func (h *MigrationHandler) Migrate() error {
	return h.DB.AutoMigrate(&models.Task{})
}
