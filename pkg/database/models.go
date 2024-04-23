package database

import "gorm.io/gorm"

// Device represents a device in the system.
type Device struct {
	gorm.Model
	Name             string `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"`
	ProductName      string `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"`
	ProviderDeviceID string `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"`
	Online           bool   `gorm:"not null;default:false" binding:"required"`
	Instruction      *string
	QrCode           *string
	SpaceID          *uint
}

// Space represents a space in the system.
type Space struct {
	gorm.Model
	Name          string `gorm:"not null;type:varchar(100);default:null" binding:"required"`
	ParentSpaceID *uint
	SubSpaces     []*Space `gorm:"foreignKey:ParentSpaceID"`
	Level         int
	Devices       []Device
	Whitelist     Whitelist
}

// User represents a user in the system.
type User struct {
	gorm.Model
	Username   string       `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"` // Username of the user
	RoleID     uint         `gorm:"not null;DEFAULT:3" json:"role_id"`
	Email      string       `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"` // Email address of the user
	Password   string       `gorm:"size:255;not null" json:"-" binding:"required"`                     // Password of the user (should be hashed)
	Name       string       `binding:"required"`                                                       // Real Name of owner of the user
	WhiteLists []*Whitelist `gorm:"many2many:user_whitelist"`                                          // many to many relation though user_whitelist table
	Role       Role         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// Log represents a log in the system.
type Log struct {
	gorm.Model
	StatusCode int
	UserID     uint
	User       User
	SpaceID    uint
	Space      Space
}

// WhiteList represents the whitelist of the device in the system.
type Whitelist struct {
	gorm.Model
	Users   []*User `gorm:"many2many:user_whitelist"`
	SpaceID uint
}

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}
