package model

import (
	"qrcheckin/internal/module/types"

	"gorm.io/gorm"
)

type Organization struct {
	_ *types.Organization
	_ *gorm.DB
}
