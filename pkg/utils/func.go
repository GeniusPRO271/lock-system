package utils

import (
	"strconv"

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

func StringToUint(str string) (uint, error) {
	// Convert the string to uint64
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err // Return 0 and the error if conversion fails
	}

	// Convert uint64 to uint
	uintNum := uint(num)

	return uintNum, nil // Return the converted uint value and nil error
}

func GetRoleNameByID(roleID uint) string {
	// Map role IDs to role names
	roles := map[uint]string{
		1: "Admin",
		2: "Verified",
		3: "Unverified",
	}

	// Retrieve role name from the map
	roleName, exists := roles[roleID]
	if !exists {
		return "unknown" // Return "unknown" if role ID doesn't exist
	}

	return roleName
}
