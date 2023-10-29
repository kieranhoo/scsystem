package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/module/entity"
)

type Organization struct {
	_ *entity.Organization
	_ *gorm.DB
}
