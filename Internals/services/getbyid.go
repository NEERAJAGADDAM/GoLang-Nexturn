package services

import (
	"net/http"
	"restcrudapi/Internals/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.Registers
	query := "SELECT * FROM registers WHERE register_id = ?"
	err = DB.QueryRow(query, id).Scan(
		&user.RegisterID,
		&user.RegisterName.FirstName,
		&user.RegisterName.LastName,
		&user.RegisterDetails.MobileNumber,
		&user.RegisterDetails.RegisterAddress.State,
		&user.RegisterDetails.RegisterAddress.City,
		&user.RegisterDetails.RegisterAddress.Zipcode,
		&user.RegisterAge,
		&user.RegisterDepartment,
		&user.RegisterSalary,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
