package space

import (
	"github.com/GeniusPRO271/lock-system/pkg/database"
)

type SpaceDTO struct {
	ID            uint
	Name          string
	ParentSpaceID *uint
	SubSpaces     []SpaceDTO
	Devices       []database.Device
}

type UpdateSpaceDTO struct {
	Name          string `json:"Name" binding:"required"`
	ParentSpaceID *uint  `json:"ParentSpaceID"`
}
