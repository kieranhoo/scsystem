package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/types/entity"
)

type Organization struct {
	_ *entity.Organization
	_ *gorm.DB
}
