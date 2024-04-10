package utils

import (
	"github.com/GeniusPRO271/lock-system/pkg/database"
	"gorm.io/gorm"
)

func PropagateWhitelist(db *gorm.DB, node *database.Space, whitelist database.Whitelist) error {
	if node == nil {
		return nil
	}

	node.Whitelist = whitelist
	if err := db.Save(&node).Error; err != nil {
		return err
	}

	for _, child := range node.SubSpaces {
		if err := PropagateWhitelist(db, child, whitelist); err != nil {
			return err
		}
	}

	return nil
}

func RemoveUserFromWhitelist(db *gorm.DB, node *database.Space, userId uint) error {
	if node == nil {
		return nil
	}

	// Remove user from node's whitelist
	var newAllowedUsers []*database.User
	for _, user := range node.Whitelist.Users {
		if user.ID != userId {
			newAllowedUsers = append(newAllowedUsers, user)
		}
	}
	node.Whitelist.Users = newAllowedUsers

	if err := db.Save(&node).Error; err != nil {
		return err
	}

	// Recursively remove user from children's whitelists
	for _, child := range node.SubSpaces {
		if err := RemoveUserFromWhitelist(db, child, userId); err != nil {
			return err
		}
	}

	return nil
}
