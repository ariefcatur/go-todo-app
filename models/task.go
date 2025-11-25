package models

import (
	"time"
)

type Task struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	UserID      int64      `gorm:"index;not null" json:"user_id"`
	Title       string     `gorm:"size:255;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Priority    string     `gorm:"size:10;default:medium;index:idx_user_priority" json:"priority"`
	Status      string     `gorm:"size:12;default:pending;index:idx_user_status" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
