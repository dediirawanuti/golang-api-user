package models

import "github.com/jinzhu/gorm"

// User structure
type User struct {
	gorm.Model
	Name     string `gorm:"column:name;size:255;not null" json:"name"`
	Email    string `gorm:"column:email;size:100;not null;unique" json:"email"`
	NoTlpn   string `gorm:"column:notlpn;not null" json:"no_tlpn"`
	Password string `gorm:"column:password;size:100;not null;" json:"password"`
	Gender   string `gorm:"column:gender;size:15;not null" json:"gender"`
	Address  string `gorm:"column:address;size:300;not null" json:"address"`
	Role     string `gorm:"column:role;default:'user'" json:"role"`
	Photo    string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

// UserWrapper use for print user data without password
type UserWrapper struct {
	gorm.Model
	Name    string `gorm:"column:name;size:255;not null" json:"name"`
	Email   string `gorm:"column:email;size:100;not null;unique" json:"email"`
	NoTlpn  string `gorm:"column:notlpn;not null" json:"no_tlpn"`
	Gender  string `gorm:"column:gender;size:15;not null" json:"gender"`
	Address string `gorm:"column:address;size:300;not null" json:"address"`
	Role    string `gorm:"column:role;default:'user'" json:"role"`
	Photo   string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

// Photo use for update photo user
type Photo struct {
	gorm.Model
	Photo string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}
