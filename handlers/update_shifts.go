package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetShifts godoc
// @Summary get shifts by dto
// @Schemes
// @Description get shifts by dto
// @Tags shifts
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get shifts by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "shifts not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /shifts [get]
func (ss *ShiftService) HandleUpdateShift(c *gin.Context) (int, interface{}, error) {

	// get shift id
	shiftID := c.Param("id")

	// get dto from req.body
	var dto models.Shift
	if err := c.ShouldBindJSON(&dto); err != nil {
		return http.StatusBadRequest, nil, errors.New("invalid shift dto")
	}

	// get shift from db
	var shift models.Shift

	// check shift exists
	if err := ss.db.Where("id = ?", shiftID).First(&shift).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnprocessableEntity, nil, errors.New("shift not found")
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	// update shift
	if err := ss.db.Model(&shift).Updates(dto).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	return http.StatusOK, shift, nil
}
