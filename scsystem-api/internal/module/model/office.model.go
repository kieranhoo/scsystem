package model

import (
	"qrcheckin/internal/module/types"

	"gorm.io/gorm"
)

type Office struct {
	_ *types.Office
	_ *gorm.DB
}
