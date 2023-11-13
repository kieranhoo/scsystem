package repo

import (
	"gorm.io/gorm"
	"qrcheckin/internal/model"
)

type Room struct {
	_ *model.Room
	_ *gorm.DB
}
