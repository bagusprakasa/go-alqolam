package handler

import (
	"go-alqolam/helper"
	"go-alqolam/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.Index()
	if err != nil {
		response := helper.APIResponse("Error to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of users", http.StatusOK, "success", user.FormatUsers(users))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Show(c *gin.Context) {
	var input user.GetDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	users, err := h.userService.Show(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get detail of user", http.StatusOK, "success", user.FormatUserNonToken(users))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Update(c *gin.Context) {
	var inputID user.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData user.UpdatedUserInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update user", http.StatusOK, "success", user.FormatUserNonToken(updatedUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Destroy(c *gin.Context) {
	var inputID user.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.Destroy(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete user", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}
