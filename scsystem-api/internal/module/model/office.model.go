package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/module/entity"
)

type Office struct {
	_ *entity.Office
	_ *gorm.DB
}
