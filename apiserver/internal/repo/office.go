package repo

import (
	"gorm.io/gorm"
	"qrcheckin/internal/model"
)

type Office struct {
	_ *model.Office
	_ *gorm.DB
}
