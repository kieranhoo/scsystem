package repo

import (
	"gorm.io/gorm"
	"qrcheckin/internal/model"
)

type Organization struct {
	_ *model.Organization
	_ *gorm.DB
}
