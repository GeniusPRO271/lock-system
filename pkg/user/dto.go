package user

// User represents a user in the system.
type Login struct {
	Email    string `json:"email" binding:"required"`    // Email address of the user
	Password string `json:"password" binding:"required"` // Password of the user (should be hashed)
}

type LoginResponse struct {
	User  UserGetResponse `json:"user" binding:"required"`
	Token string          `json:"token" binding:"required"`
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
	Id       uint   `json:"id"`
	Username string `json:"username"` // Username of the user
	Email    string `json:"email"`    // Password of the user (should be hashed)
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UsersGetResponse struct {
	Users []UserGetResponse
}
