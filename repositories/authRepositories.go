package repositories

import (
	"errors"
	"management-book/configs"
	"management-book/models"

	"golang.org/x/crypto/bcrypt"
)

func Login(userLogin *models.UserLogin, user *models.User) error {
	var checkUser models.User

	userData := configs.DB.Where("email = ?", userLogin.Email).First(&checkUser)

	// check user data
	if userData.Error != nil {
		return errors.New("akun tidak terdaftar, silahkan daftar terlebih dahulu")
	}

	// check the user data password
	checkPassword := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(userLogin.Password))

	// check password
	if checkPassword != nil {
		return errors.New("password yang dimasukkan salah, silahkan coba lagi")
	}

	// get all users
	configs.DB.First(&user)

	return nil
}

func Register(user *models.User) error {
	result := configs.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
