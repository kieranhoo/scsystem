package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/types/entity"
)

type Office struct {
	_ *entity.Office
	_ *gorm.DB
}
