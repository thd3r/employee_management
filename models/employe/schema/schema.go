package schema

type CreateEmployeSchema struct {
	Name string `json:"name" validate:"required,min=5,max=50"`
	Role string `json:"role" validate:"required,min=4,max=20"`
}

type UpdateEmployeSchema struct {
	Name string `json:"name" validate:"required,min=5,max=50"`
	Role string `json:"role" validate:"required,min=4,max=20"`
}
