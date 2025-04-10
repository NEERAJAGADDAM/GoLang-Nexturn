package services

import (
	"net/http"
	"restcrudapi/Internals/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser models.Registers
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	query := `
        UPDATE registers 
        SET first_name=?, last_name=?, mobile_number=?, state=?, city=?, zipcode=?, age=?, department=?, salary=?
        WHERE register_id=?
    `
	_, err = DB.Exec(query,
		updatedUser.RegisterName.FirstName,
		updatedUser.RegisterName.LastName,
		updatedUser.RegisterDetails.MobileNumber,
		updatedUser.RegisterDetails.RegisterAddress.State,
		updatedUser.RegisterDetails.RegisterAddress.City,
		updatedUser.RegisterDetails.RegisterAddress.Zipcode,
		updatedUser.RegisterAge,
		updatedUser.RegisterDepartment,
		updatedUser.RegisterSalary,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
