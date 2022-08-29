package handler

import (
	"go-alqolam/helper"
	"go-alqolam/region"
	"net/http"

	"github.com/gin-gonic/gin"
)

type regionHandler struct {
	regionService region.Service
}

func NewRegionHandler(regionService region.Service) *regionHandler {
	return &regionHandler{regionService}
}

func (h *regionHandler) Index(c *gin.Context) {
	regions, err := h.regionService.Index()
	if err != nil {
		response := helper.APIResponse("Error to get regions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of regions", http.StatusOK, "success", region.FormatRegions(regions))
	c.JSON(http.StatusOK, response)
}

func (h *regionHandler) Store(c *gin.Context) {
	var input region.RegionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Store region failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRegion, err := h.regionService.Store(input)

	if err != nil {
		response := helper.APIResponse("Store region failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Store region failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := region.FormatRegion(newRegion)

	response := helper.APIResponse("Success store region", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *regionHandler) Show(c *gin.Context) {
	var input region.GetDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of region", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	regions, err := h.regionService.Show(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of region", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get detail of region", http.StatusOK, "success", region.FormatRegion(regions))
	c.JSON(http.StatusOK, response)
}

func (h *regionHandler) Update(c *gin.Context) {
	var inputID region.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update region", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData region.RegionInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update region", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedRegion, err := h.regionService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update region", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update region", http.StatusOK, "success", region.FormatRegion(updatedRegion))
	c.JSON(http.StatusOK, response)
}

func (h *regionHandler) Destroy(c *gin.Context) {
	var inputID region.GetDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete region", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	region, err := h.regionService.Destroy(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update region", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete region", http.StatusOK, "success", region)
	c.JSON(http.StatusOK, response)
}
