package database

import "gorm.io/gorm"

// Device represents a device in the system.
type Device struct {
	gorm.Model
	Name        string `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"`
	DeviceModel string `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"`
	Type        string `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"`
	Instruction string
	QrCode      string
}

// User represents a user in the system.
type User struct {
	gorm.Model
	Username   string       `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"` // Username of the user
	Email      string       `gorm:"unique;not null;type:varchar(100);default:null" binding:"required"` // Email address of the user
	Password   string       `binding:"required"`                                                       // Password of the user (should be hashed)
	Name       string       `binding:"required"`                                                       // Real Name of owner of the user
	Log        Log          // LogID from the log
	WhiteLists []*Whitelist `gorm:"many2many:user_whitelist;"` // many to many relation though user_whitelist table
}

type Log struct {
	gorm.Model
	UserID uint
}

// WhiteList represents the whitelist of the device in the system.
type Whitelist struct {
	gorm.Model
	Users    []*User `gorm:"many2many:user_whitelist;"`
	DeviceID uint
	Device   Device
}
