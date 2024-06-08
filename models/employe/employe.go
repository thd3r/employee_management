package employe

import "time"

type Employe struct {
	Id        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(250)" json:"name"`
	Role      string    `gorm:"type:varchar(50)" json:"role"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}
