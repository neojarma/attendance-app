package model

type Employees struct {
	Nip         string `json:"nip"`
	Name        string `json:"name"`
	Position_Id string `json:"positionId"`
	Role_Id     string `json:"roleId"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Is_Active   bool   `json:"isActive"`
}

type EmployeesJoinRole struct {
	Nip         string `json:"nip"`
	Name        string `json:"name"`
	Position_Id string `json:"positionId"`
	Role        string `json:"role"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Is_Active   bool   `json:"isActive"`
}
