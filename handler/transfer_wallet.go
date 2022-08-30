package handler

import (
	"go-alqolam/helper"
	transferWallet "go-alqolam/transfer_wallet"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transferWalletHandler struct {
	transferWalletService transferWallet.Service
}

func NewTransferWalletHandler(transferWalletService transferWallet.Service) *transferWalletHandler {
	return &transferWalletHandler{transferWalletService}
}

func (h *transferWalletHandler) Index(c *gin.Context) {
	transferWallets, err := h.transferWalletService.Index()
	if err != nil {
		response := helper.APIResponse("Error to get transfer wallets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of transfer wallets", http.StatusOK, "success", transferWallet.FormatTransferWallets(transferWallets))
	c.JSON(http.StatusOK, response)
}

func (h *transferWalletHandler) Store(c *gin.Context) {
	var input transferWallet.TransferWalletInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Store transfer wallet failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTransferWallet, err := h.transferWalletService.Store(input)

	if err != nil {
		response := helper.APIResponse("Store transfer wallet failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Store transfer wallet failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transferWallet.FormatTransferWallet(newTransferWallet)

	response := helper.APIResponse("Success wallet region", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *transferWalletHandler) Show(c *gin.Context) {
	var input transferWallet.GetDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of transfer wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	transferWallets, err := h.transferWalletService.Show(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of transfer wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get detail of wallet", http.StatusOK, "success", transferWallet.FormatTransferWallet(transferWallets))
	c.JSON(http.StatusOK, response)
}

func (h *transferWalletHandler) Update(c *gin.Context) {
	var inputID transferWallet.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update transfer wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData transferWallet.TransferWalletInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update transfer wallet", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedTransferWallet, err := h.transferWalletService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update transfer wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update transfer wallet", http.StatusOK, "success", transferWallet.FormatTransferWallet(updatedTransferWallet))
	c.JSON(http.StatusOK, response)
}

func (h *transferWalletHandler) Destroy(c *gin.Context) {
	var inputID transferWallet.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete transfer wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	wallet, err := h.transferWalletService.Destroy(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update transfer wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete transfer wallet", http.StatusOK, "success", wallet)
	c.JSON(http.StatusOK, response)
}
