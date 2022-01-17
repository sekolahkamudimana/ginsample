package handler

import (
	"ginsample/helper"
	"ginsample/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//map input dari user ke struct register
	//struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account has failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		log.Println(response)
		return

		// info := `#R maaf gagal.`
		// c.String(http.StatusAccepted, info)
		// return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account has failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	// token,err := h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "token1234567890")

	response := helper.ApiResponse("Account has been registered", http.StatusAccepted, "success", formatter)

	c.JSON(http.StatusOK, response)
}
