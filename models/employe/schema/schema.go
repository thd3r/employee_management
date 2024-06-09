package schema

type CreateEmployeSchema struct {
	Name 	 string  `json:"name" validate:"required,min=5,required,max=50"`
	Role 	 string  `json:"role" validate:"required,min=4,required,max=20"`
	Salary 	 string  `json:"salary" validate:"required,max=200"`
	Position string  `json:"position" validate:"required,max=100"`
}

type UpdateEmployeSchema struct {
	Name 	 string  `json:"name" validate:"max=50"`
	Role	 string  `json:"role" validate:"max=20"`
	Salary 	 string  `json:"salary" validate:"max=200"`
	Position string  `json:"position" validate:"max=100"`
}
