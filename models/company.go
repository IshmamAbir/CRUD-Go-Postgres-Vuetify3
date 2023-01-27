package models

import (
	"github.com/google/uuid"
)

// Company Model
type Company struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	Location string    `json:"location" gorm:"type:varchar(200);not null"`
}

// this method set the table name when gorm creates the table
func (Company) TableName() string {
	return "company"
}
