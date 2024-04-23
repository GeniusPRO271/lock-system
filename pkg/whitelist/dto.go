package whitelist

import "github.com/GeniusPRO271/lock-system/pkg/device"

type PostDeleteWhiteListParams struct {
	SpaceID   uint  `json:"SpaceID" binding:"required"`
	UserID    uint  `json:"UserID" binding:"required"`
	Propagate *bool `json:"Propagate" binding:"required"`
}

type SpaceWhitelistDTO struct {
	ID            uint
	Name          string
	ParentSpaceID *uint
	SubSpaces     []SpaceWhitelistDTO
	Devices       []device.DevicesGetResponse
	IsUserAllowed *bool
}
