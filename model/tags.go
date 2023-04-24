package model

type Tags struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Slug        string `gorm:"type:varchar(255);" json:"slug"`
	Type        string `gorm:"type:varchar(255);" json:"type"`
	Category    string `gorm:"type:varchar(255);" json:"category"`
	Status      string `gorm:"type:varchar(255);" json:"status"`
}
