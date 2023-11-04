package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/types"
)

type Organization struct {
	_ *types.Organization
	_ *gorm.DB
}
