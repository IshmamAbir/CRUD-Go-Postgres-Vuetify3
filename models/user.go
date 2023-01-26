package models

import "github.com/google/uuid"

// user Model
type User struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"firstName" gorm:"not null;type:varchar(100)"`
	LastName  string    `json:"lastName" gorm:"not null;type:varchar(100)"`
	Email     string    `json:"email" gorm:"not null;type:varchar(100)"`
}
