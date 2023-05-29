package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
func (ss *ShiftService) HandleGetShift(c *gin.Context) (int, interface{}, error) {

	// get shifts
	db, err := sql.Open("postgres", "postgres://shiftuser:shiftdb@localhost:5432/shiftdb?sslmode=disable")
	if err != nil {
		log.Print("err")
		return http.StatusInternalServerError, nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM shifts")
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, nil, err
	}
	defer rows.Close()
	var data []models.Shift
	log.Print("err1")
	for rows.Next() {
		var shift models.Shift
		err := rows.Scan(
			&shift.ID,
			&shift.ShiftID,
			&shift.UserID,
			&shift.StartTime,
			&shift.EndTime,
			&shift.CreatedAt,
			&shift.UpdatedAt,
			&shift.MadeField,
		)
		log.Print("err2")
		if err != nil {
			log.Print(err)
			return http.StatusInternalServerError, nil, err
		}
		log.Print("err3")
		data = append(data, shift) // shifti slice'a ekle
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
		return http.StatusInternalServerError, nil, err
	}
	log.Print("err4")

	return http.StatusOK, data, nil

}
