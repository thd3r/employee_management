package schema

type CreateEmployeeSchema struct {
	Name     string `json:"name" validate:"required,min=5,required,max=50"`
	Salary   string `json:"salary" validate:"required,max=200"`
	Position string `json:"position" validate:"required,max=100"`
}

type UpdateEmployeeSchema struct {
	Name     string `json:"name" validate:"max=50"`
	Salary   string `json:"salary" validate:"max=200"`
	Position string `json:"position" validate:"max=100"`
}
