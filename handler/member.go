package handler

import (
	"go-alqolam/helper"
	"go-alqolam/member"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberHandler struct {
	memberService member.Service
}

func NewMemberHandler(memberService member.Service) *memberHandler {
	return &memberHandler{memberService}
}

func (h *memberHandler) Index(c *gin.Context) {
	members, err := h.memberService.Index()
	if err != nil {
		response := helper.APIResponse("Error to get members", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of members", http.StatusOK, "success", member.FormatMembers(members))
	c.JSON(http.StatusOK, response)
}

func (h *memberHandler) Store(c *gin.Context) {
	var input member.MemberInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Store member failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMember, err := h.memberService.Store(input)

	if err != nil {
		response := helper.APIResponse("Store member failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Store member failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := member.FormatMember(newMember)

	response := helper.APIResponse("Success store member", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *memberHandler) Show(c *gin.Context) {
	var input member.GetDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of member", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	members, err := h.memberService.Show(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of member", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get detail of member", http.StatusOK, "success", member.FormatMember(members))
	c.JSON(http.StatusOK, response)
}

func (h *memberHandler) Update(c *gin.Context) {
	var inputID member.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update member", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData member.MemberInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update member", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedMember, err := h.memberService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update member", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update member", http.StatusOK, "success", member.FormatMember(updatedMember))
	c.JSON(http.StatusOK, response)
}

func (h *memberHandler) Destroy(c *gin.Context) {
	var inputID member.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete member", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	member, err := h.memberService.Destroy(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update member", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete member", http.StatusOK, "success", member)
	c.JSON(http.StatusOK, response)
}
