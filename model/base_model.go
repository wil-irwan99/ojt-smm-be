package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"createdAt"`
}
