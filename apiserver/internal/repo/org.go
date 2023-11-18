package repo

import (
	"scsystem/internal/model"

	"gorm.io/gorm"
)

type Organization struct {
	_ *model.Organization
	_ *gorm.DB
}
