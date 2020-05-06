package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Device :Device DB
type Device struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Address   string    `gorm:"not null" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Device) SaveDevice(db *gorm.DB) (*Device, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Device{}, err
	}
	return u, nil
}
