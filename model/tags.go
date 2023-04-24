package model

type Tags struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
