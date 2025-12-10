package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID        string         `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	Email       string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string         `gorm:"type:varchar(255);not null" json:"-"`
	NamaDepan   string         `gorm:"type:varchar(100)" json:"nama_depan"`
	NamaBelakang string        `gorm:"type:varchar(100)" json:"nama_belakang"`
	ImageURL    *string        `gorm:"type:text" json:"image_url"`

	Role        string         `gorm:"type:varchar(20);not null;check:role IN ('admin','owner')" json:"role"`
	Status      string         `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','inactive','banned')" json:"status"`

	CreatedAt   time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	ModifiedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"modified_at"`

	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
