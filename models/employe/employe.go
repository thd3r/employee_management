package employe

import (
	"time"
)

type Employe struct {
	Id        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(250)" json:"name"`
	Role      string    `gorm:"type:varchar(50)" json:"role"`
	Salary    string    `gorm:"type:text; not null" json:"salary"`
	Position  string    `gorm:"type:text; not null" json:"position"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}
