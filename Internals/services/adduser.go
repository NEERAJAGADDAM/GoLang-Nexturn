package services

import (
	"net/http"
	"restcrudapi/Internals/models"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newRegister models.Registers
	if err := c.ShouldBindJSON(&newRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	query := `
        INSERT INTO registers (first_name, last_name, mobile_number, state, city, zipcode, age, department, salary)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	result, err := DB.Exec(query,
		newRegister.RegisterName.FirstName,
		newRegister.RegisterName.LastName,
		newRegister.RegisterDetails.MobileNumber,
		newRegister.RegisterDetails.RegisterAddress.State,
		newRegister.RegisterDetails.RegisterAddress.City,
		newRegister.RegisterDetails.RegisterAddress.Zipcode,
		newRegister.RegisterAge,
		newRegister.RegisterDepartment,
		newRegister.RegisterSalary,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	id, _ := result.LastInsertId()
	newRegister.RegisterID = int(id)

	c.JSON(http.StatusCreated, newRegister)
}
