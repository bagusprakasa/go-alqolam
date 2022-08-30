package handler

import (
	"go-alqolam/helper"
	"go-alqolam/region"
	"go-alqolam/wallet"
	"net/http"

	"github.com/gin-gonic/gin"
)

type walletHandler struct {
	walletService wallet.Service
}

func NewWalletHandler(walletService wallet.Service) *walletHandler {
	return &walletHandler{walletService}
}

func (h *walletHandler) Index(c *gin.Context) {
	wallets, err := h.walletService.Index()
	if err != nil {
		response := helper.APIResponse("Error to get wallets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of wallets", http.StatusOK, "success", wallet.FormatWallets(wallets))
	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) Store(c *gin.Context) {
	var input wallet.WalletInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Store wallet failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newWallet, err := h.walletService.Store(input)

	if err != nil {
		response := helper.APIResponse("Store wallet failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Store wallet failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := wallet.FormatWallet(newWallet)

	response := helper.APIResponse("Success wallet region", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) Show(c *gin.Context) {
	var input region.GetDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	wallets, err := h.walletService.Show(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get detail of wallet", http.StatusOK, "success", wallet.FormatWallet(wallets))
	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) Update(c *gin.Context) {
	var inputID wallet.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData wallet.WalletInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update wallet", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedWallet, err := h.walletService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update wallet", http.StatusOK, "success", wallet.FormatWallet(updatedWallet))
	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) Destroy(c *gin.Context) {
	var inputID wallet.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	wallet, err := h.walletService.Destroy(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update wallet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete wallet", http.StatusOK, "success", wallet)
	c.JSON(http.StatusOK, response)
}
