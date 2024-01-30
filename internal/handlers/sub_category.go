package handlers

import (
	"FinanceTODO/internal/handlers/responses"
	"FinanceTODO/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateSubCategory(c *gin.Context) {
	userID := getUserID(c)

	var input models.SubCategoryInput

	if err := c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	id, err := h.services.SubCategory.Create(input, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetSubCategories(c *gin.Context) {
	userID := getUserID(c)

	subCategories, err := h.services.SubCategory.GetAll(userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, subCategories)
}

func (h *Handler) GetSubCategoryById(c *gin.Context) {
	userID := getUserID(c)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	subCategory, err := h.services.SubCategory.GetByID(id, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, subCategory)
}

func (h *Handler) UpdateSubCategory(c *gin.Context) {
	userID := getUserID(c)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var input models.SubCategoryInput

	if err = c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	err = h.services.SubCategory.Update(input, id, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteSubCategory(c *gin.Context) {
	userID := getUserID(c)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	err = h.services.SubCategory.Delete(id, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, map[string]interface{}{
		"status": "ok",
	})
}
