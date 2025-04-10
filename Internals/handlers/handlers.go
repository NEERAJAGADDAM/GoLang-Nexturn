package handlers

import "restcrudapi/Internals/models"

var Registers = []models.Registers{
	{
		RegisterID: 1,
		RegisterName: models.RegisterName{
			FirstName: "Gaddam",
			LastName:  "Neeraja",
		},
		RegisterDetails: models.RegisterDetails{
			MobileNumber: "9392841634",
			RegisterAddress: models.RegisterAddress{
				City:    "Warangal",
				State:   "Telangana",
				Zipcode: "506366",
			},
		},
		RegisterAge:        21,
		RegisterDepartment: "Engineering",
		RegisterSalary:     50000,
	},
}
