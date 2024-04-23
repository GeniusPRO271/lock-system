package whitelist

import (
	"errors"
	"strconv"

	"github.com/GeniusPRO271/lock-system/pkg/database"
	deviceModel "github.com/GeniusPRO271/lock-system/pkg/device"
	spacepkg "github.com/GeniusPRO271/lock-system/pkg/space"
	"github.com/GeniusPRO271/lock-system/pkg/user"
	"github.com/GeniusPRO271/lock-system/pkg/utils"
	"gorm.io/gorm"
)

type WhitelistService interface {
	AddUserToWhitelist(userId, spaceId uint, propagate bool) error
	GetUsersFromSpaceWhitelist(spaceID string) ([]user.UserGetResponse, error)
	DeleteUserFromWhitelist(userId, spaceId uint, propagate bool) error
	IsUserInWhitelist(userId uint, spaceId uint) (bool, error)
	GetRootSpaces(userID uint) ([]SpaceWhitelistDTO, error)
	SpaceToWhitelistDTORecursive(spaceData database.Space, userID uint) SpaceWhitelistDTO
}

type WhitelistServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	SpaceService spacepkg.SpaceService
	Db           *gorm.DB
}

func (s *WhitelistServiceImpl) GetRootSpaces(userID uint) ([]SpaceWhitelistDTO, error) {
	var spaces []database.Space
	// Preload the whitelist
	if err := s.Db.Preload("Whitelist.Users").Preload("Devices").Where("level = ?", 1).Find(&spaces).Error; err != nil {
		return nil, err
	}

	var fullSpacelist []SpaceWhitelistDTO
	for _, space := range spaces {
		if err := s.SpaceService.LoadSubSpaces(&space); err != nil {
			return nil, err
		}
		spaceDTO := s.SpaceToWhitelistDTORecursive(space, userID) // Use the correct recursive function
		fullSpacelist = append(fullSpacelist, spaceDTO)
	}
	return fullSpacelist, nil
}

func (s *WhitelistServiceImpl) IsUserInWhitelist(userId uint, spaceId uint) (bool, error) {
	var user database.User
	if err := s.Db.Preload("WhiteLists").First(&user, userId).Error; err != nil {
		return false, err
	}

	// Check if the user has any whitelists
	for _, whitelist := range user.WhiteLists {
		if whitelist.SpaceID == spaceId {
			return true, nil
		}
	}

	// If no whitelist contains the specified space, return false
	return false, nil
}

func (s *WhitelistServiceImpl) AddUserToWhitelist(userId, spaceId uint, propagate bool) error {

	var space database.Space
	if err := s.Db.Preload("Whitelist").Preload("SubSpaces").First(&space, spaceId).Error; err != nil {
		return err
	}

	var user database.User
	if err := s.Db.First(&user, userId).Error; err != nil {
		return err
	}

	// Add the user to the whitelist
	space.Whitelist.Users = append(space.Whitelist.Users, &user)

	// Save the space, which will also save the associated whitelist
	if err := s.Db.Save(&space).Error; err != nil {
		return err
	}

	// Recursively add the user to the whitelists of subspaces if propagate is true
	if propagate {
		for i := range space.SubSpaces {
			if err := s.AddUserToWhitelist(userId, space.SubSpaces[i].ID, propagate); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *WhitelistServiceImpl) GetUsersFromSpaceWhitelist(spaceId string) ([]user.UserGetResponse, error) {
	spaceID, err := strconv.ParseUint(spaceId, 10, 64)
	if err != nil {
		return nil, err
	}

	var space database.Space
	if err := s.Db.Preload("Whitelist.Users").First(&space, spaceID).Error; err != nil {
		return nil, err
	}

	var users []user.UserGetResponse
	for _, u := range space.Whitelist.Users {
		user := user.UserGetResponse{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Name:     u.Name,
			Role:     utils.GetRoleNameByID(u.RoleID),
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *WhitelistServiceImpl) DeleteUserFromWhitelist(userId, spaceId uint, propagate bool) error {
	// Find the whitelist for the specified space ID
	var whitelist database.Whitelist
	if err := s.Db.Preload("Users").First(&whitelist, "space_id = ?", spaceId).Error; err != nil {
		return err
	}

	// Find the user in the whitelist
	var userToDelete *database.User
	for _, user := range whitelist.Users {
		if user.ID == userId {
			userToDelete = user
			break
		}
	}

	if userToDelete == nil {
		return errors.New("user not found in the whitelist")
	}

	// Use GORM's association to delete the user from the whitelist
	if err := s.Db.Model(&whitelist).Association("Users").Delete(userToDelete); err != nil {
		return err
	}

	// If propagate is true, recursively delete the user from the whitelists of subspaces
	if propagate {
		var space database.Space
		if err := s.Db.Preload("SubSpaces").First(&space, spaceId).Error; err != nil {
			return err
		}

		for _, subSpace := range space.SubSpaces {
			if err := s.DeleteUserFromWhitelist(userId, subSpace.ID, propagate); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *WhitelistServiceImpl) SpaceToWhitelistDTORecursive(spaceData database.Space, userID uint) SpaceWhitelistDTO {
	var subSpacesDTO []SpaceWhitelistDTO
	for _, subSpace := range spaceData.SubSpaces {
		subSpaceDTO := s.SpaceToWhitelistDTORecursive(*subSpace, userID)
		subSpacesDTO = append(subSpacesDTO, subSpaceDTO)
	}

	var isUserAllowed = false
	for _, user := range spaceData.Whitelist.Users {
		if user.ID == userID {
			allowed := true
			isUserAllowed = allowed
			break
		}
	}

	var devices []deviceModel.DevicesGetResponse
	for _, device := range spaceData.Devices {
		devices = append(devices, deviceModel.DevicesGetResponse{
			Id:               device.ID,
			Online:           device.Online,
			Name:             device.Name,
			ProductName:      device.ProductName,
			ProviderDeviceID: device.ProviderDeviceID,
			SpaceID:          device.SpaceID,
		})
	}

	spaceWhitelistDTO := SpaceWhitelistDTO{
		ID:            spaceData.ID,
		Name:          spaceData.Name,
		ParentSpaceID: spaceData.ParentSpaceID,
		SubSpaces:     subSpacesDTO,
		Devices:       devices,
		IsUserAllowed: &isUserAllowed,
	}

	return spaceWhitelistDTO
}
