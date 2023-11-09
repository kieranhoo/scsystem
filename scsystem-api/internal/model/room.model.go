package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/types"
)

type Room struct {
	_ *types.Room
	_ *gorm.DB
}
