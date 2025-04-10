package models

type RegisterName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegisterAddress struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type RegisterDetails struct {
	MobileNumber    string          `json:"mobile_number"`
	RegisterAddress RegisterAddress `json:"register_address"`
}

type Registers struct {
	RegisterID         int             `json:"register_id"`
	RegisterName       RegisterName    `json:"register_name"`
	RegisterDetails    RegisterDetails `json:"register_details"`
	RegisterAge        int             `json:"register_age"`
	RegisterDepartment string          `json:"register_department"`
	RegisterSalary     float64         `json:"register_salary"`
}
