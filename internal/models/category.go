package models

type Category struct {
	ID     uint     `gorm:"primaryKey"`
	Name   string   `gorm:"unique;not null"`
	Movies []*Movie `gorm:"many2many:movie_categories;"`
}
