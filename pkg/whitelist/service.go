package whitelist

import (
	"errors"
	"strconv"

	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/user"
	"gorm.io/gorm"
)

type WhitelistService interface {
	AddUserToWhitelist(userId, deviceId string) error
	GetUsersFromSpaceWhitelist(spaceID string) ([]user.UserGetResponse, error)
	DeleteUserFromWhitelist(userId, spaceId string) error
}

type WhitelistServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *WhitelistServiceImpl) AddUserToWhitelist(userId, spaceId string) error {
	spaceID, err := strconv.ParseUint(spaceId, 10, 64)
	if err != nil {
		return err
	}

	var space database.Space
	if err := s.Db.Preload("Whitelist").First(&space, spaceID).Error; err != nil {
		return err
	}

	userID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return err
	}

	var user database.User
	if err := s.Db.First(&user, userID).Error; err != nil {
		return err
	}

	// Add the user to the whitelist
	space.Whitelist.Users = append(space.Whitelist.Users, &user)

	// Save the space, which will also save the associated whitelist
	if err := s.Db.Save(&space).Error; err != nil {
		return err
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
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *WhitelistServiceImpl) DeleteUserFromWhitelist(userId, spaceId string) error {
	spaceID, err := strconv.ParseUint(spaceId, 10, 64)
	if err != nil {
		return err
	}

	userID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return err
	}

	// Find the whitelist for the specified space ID
	var whitelist database.Whitelist
	if err := s.Db.Preload("Users").First(&whitelist, "space_id = ?", spaceID).Error; err != nil {
		return err
	}

	// Find the user in the whitelist
	var userToDelete database.User
	for _, user := range whitelist.Users {
		if user.ID == uint(userID) {
			userToDelete = *user
			break
		}
	}

	if userToDelete.ID == 0 {
		return errors.New("user not found in the whitelist")
	}

	// Use GORM's association to delete the user from the whitelist
	if err := s.Db.Model(&whitelist).Association("Users").Delete(&userToDelete); err != nil {
		return err
	}

	return nil
}
