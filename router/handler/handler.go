package handler

import (
	"github.com/shrtkv/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

// InitializeHandler ...
func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
