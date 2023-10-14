package controllers

import (
	"management-book/middlewares"
	"management-book/models"
	"management-book/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c echo.Context) error {
	var user models.User
	var userLogin models.UserLogin

	c.Bind(&userLogin)

	// error validation form
	errValidate := c.Validate(userLogin)

	// check the validation
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// Login process
	errLogin := repositories.Login(&userLogin, &user)

	// check the error
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errLogin.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// make a new response
	var userLoginResponse models.UserLoginResponse

	// append to a new response
	userLoginResponse.Id = user.Id
	userLoginResponse.Name = user.Name
	userLoginResponse.Email = user.Email
	userLoginResponse.Token = middlewares.GenerateJWTToken(userLoginResponse.Id, user.Name)

	// return success response
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Login berhasil",
		Status:  true,
		Data:    userLoginResponse,
	})
}

func RegisterController(c echo.Context) error {
	var userRegister models.User

	c.Bind(&userRegister)

	// error validation form
	errValidate := c.Validate(&userRegister)

	// check the validation
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// make bcrypt password
	result, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)

	// check the error
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// change []byte to string
	userRegister.Password = string(result)

	// register process
	errRegister := repositories.Register(&userRegister)

	// check the error
	if errRegister != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errRegister.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// make a new response
	var userResponse models.UserRegisterResponse

	// append to a new response
	userResponse.Id = userRegister.Id
	userResponse.Name = userRegister.Name
	userResponse.Email = userRegister.Email

	// return success response
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "User telah berhasil didaftarkan",
		Status:  true,
		Data:    userResponse,
	})
}
