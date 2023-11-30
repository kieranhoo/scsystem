package repo

import (
	"scsystem/internal/model"

	"gorm.io/gorm"
)

type Room struct {
	_ *model.Room
	_ *gorm.DB
}
