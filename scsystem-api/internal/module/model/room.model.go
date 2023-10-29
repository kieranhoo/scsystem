package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/module/entity"
)

type Room struct {
	_ *entity.Room
	_ *gorm.DB
}
