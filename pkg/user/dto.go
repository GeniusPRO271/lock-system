package user

// User represents a user in the system.
type Login struct {
	Email    string `json:"email" binding:"required"`    // Email address of the user
	Password string `json:"password" binding:"required"` // Password of the user (should be hashed)
}

type Register struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Update struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
}

type UserGetResponse struct {
	Id       uint
	Username string // Username of the user
	Email    string // Password of the user (should be hashed)
	Name     string
	Role     uint
}

type UsersGetResponse struct {
	Users []UserGetResponse
}
