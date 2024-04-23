package user

import (
	"log"

	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/jwt"
	"github.com/GeniusPRO271/lock-system/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user Register) error
	GetUser(User *model.User, id int) (err error)
	GetUsers() (*UsersGetResponse, error)
	VerifyUser(userCredentials Login) (*string, error)
	UpdateUser(User *model.User) (err error)
}

type UserServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *UserServiceImpl) CreateUser(user Register) error {
	// Implement registration logic here.

	hashPass, err := hashPassword(user.Password)

	if err != nil {
		return err
	}

	userDB := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashPass,
		Username: user.Username,
	}

	if result := s.Db.Create(&userDB); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserServiceImpl) VerifyUser(userCredentials Login) (*string, error) {
	var user model.User

	if err := s.Db.First(&user, "email = ?", userCredentials.Email).Error; err != nil {
		return nil, err
	}

	if err := ValidateUserPassword(&user, userCredentials.Password); err != nil {
		return nil, err
	}
	log.Printf("Passed Password")

	log.Printf("Passed DB")
	token, err := jwt.GenerateJWT(user)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserServiceImpl) GetUser(User *model.User, id int) (err error) {
	err = s.Db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
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
			Role:     utils.GetRoleNameByID(user.RoleID),
		})
	}

	return response, nil
}

func (s *UserServiceImpl) UpdateUser(user *model.User) error {
	// Add a WHERE condition to specify which user to update
	err := s.Db.Omit("password").Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	// Generate a salted hash for the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidateUserPassword(user *model.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
