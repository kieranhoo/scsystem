package repo

import (
	"scsystem/internal/model"

	"gorm.io/gorm"
)

type Office struct {
	_ *model.Office
	_ *gorm.DB
}
