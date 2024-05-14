package modules

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"not null" validate:"required,alpha,min=2,max=100"`
	LastName  string `json:"lastName" gorm:"not null" validate:"required,alpha,min=2,max=100"`
	Email     string `json:"email" gorm:"not null;uniqueIndex" validate:"required,email"`
	Password  string `json:"password" gorm:"not null" validate:"required,min=10,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`
}

type Session struct {
	AccountID    uint   `gorm:"foreignKey;not null;index;constraint:OnDelete:CASCADE;"`
	RefreshToken string `gorm:"primaryKey;size:512"`
	ExpiresAt    int64  `gorm:"index"`
}

type LoginCredentials struct {
	Email    string `json:"email" gorm:"not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,min=8"`
}
