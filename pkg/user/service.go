package user

import (
	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user model.User) error
	GetUserByID(id string) (*UserGetResponse, error)
	GetUsers() (*UsersGetResponse, error)
	VerifyUser(userCredentials UserLogin) (*string, error)
}

type UserServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *UserServiceImpl) CreateUser(user model.User) error {
	// Implement registration logic here.

	hashPass, err := hashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashPass
	if result := s.Db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserServiceImpl) VerifyUser(userCredentials UserLogin) (*string, error) {
	var user model.User

	if err := s.Db.First(&user, "email = ?", userCredentials.Email).Error; err != nil {
		return nil, err
	}

	if err := verifyPassword(user.Password, userCredentials.Password); err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(user.ID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserServiceImpl) GetUserByID(id string) (*UserGetResponse, error) {
	var user model.User

	if err := s.Db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &UserGetResponse{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *UserServiceImpl) GetUsers() (*UsersGetResponse, error) {
	var users []*model.User

	// Fetch all users from the database
	if err := s.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	// Prepare the response
	response := &UsersGetResponse{}

	// Populate the response slice
	for _, user := range users {
		response.Users = append(response.Users, UserGetResponse{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Name:     user.Name,
		})
	}

	return response, nil
}

func hashPassword(password string) (string, error) {
	// Generate a salted hash for the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, password string) error {
	// Compare the hashed password with the password provided by the user.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
