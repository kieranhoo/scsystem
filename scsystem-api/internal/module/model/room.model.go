package model

import (
	"qrcheckin/internal/module/types"

	"gorm.io/gorm"
)

type Room struct {
	_ *types.Room
	_ *gorm.DB
}
