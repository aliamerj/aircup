package modules

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}
type Disk struct {
	DiskPath string `json:"disk_path" gorm:"not null"`
}
