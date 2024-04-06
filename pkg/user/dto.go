package user

// User represents a user in the system.
type UserLogin struct {
	Email    string `binding:"required"` // Email address of the user
	Password string `binding:"required"` // Password of the user (should be hashed)
}

type UserGetResponse struct {
	Id       uint
	Username string // Username of the user
	Email    string // Password of the user (should be hashed)
	Name     string
}

type UsersGetResponse struct {
	Users []UserGetResponse
}
