package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/types"
)

type Office struct {
	_ *types.Office
	_ *gorm.DB
}
