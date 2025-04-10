package services

import (
	"net/http"
	"restcrudapi/Internals/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM registers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []models.Registers

	for rows.Next() {
		var user models.Registers
		err := rows.Scan(
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading user data"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
