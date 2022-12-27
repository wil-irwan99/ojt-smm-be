package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time      `json:"updatedAt"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
