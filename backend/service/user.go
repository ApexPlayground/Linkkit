package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(user model.User) (model.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return model.User{}, fmt.Errorf("name, email, and password are required")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password")
	}

	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return model.User{}, fmt.Errorf("email already exists")
		}
		return model.User{}, fmt.Errorf("failed to create user")
	}
	userResponse := model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
	}

	return userResponse, nil

}

func Login(email, password string) (model.User, error) {
	if email == "" || password == "" {
		return model.User{}, fmt.Errorf("email, and password are required")
	}

	var user model.User
	result := config.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return model.User{}, fmt.Errorf("invalid email or password")

	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("invalid email or password")
	}

	return user, nil

}

func GetUserByID(userID uint) (model.User, error) {
	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func ListUsers() ([]model.User, error) {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

type UpdateUserInput struct {
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

func UpdateUser(userID uint, input UpdateUserInput) (model.User, error) {
	var user model.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		return model.User{}, errors.New("user not found")
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, errors.New("failed to hash password")
		}
		user.Password = string(hashed)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return model.User{}, errors.New("failed to update user")
	}

	return user, nil
}

func DeleteUser(userID uint) error {
	var user model.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
