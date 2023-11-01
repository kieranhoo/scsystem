package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/types/entity"
)

type Room struct {
	_ *entity.Room
	_ *gorm.DB
}
