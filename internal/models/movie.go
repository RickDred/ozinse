package models

type Movie struct {
	ID          uint `gorm:"primarykey"`
	Logo        string
	Title       string
	Description string
	Producer    string
	Director    string
	Year        string
	Type        string
	Seasons     int
	Episodes    int
	Minutes     int
	Screenshots []string `gorm:"type:text[]"`
}

type Video struct {
	ID       uint `gorm:"primarykey"`
	MovieID  uint
	Movie    Movie `gorm:"constraint:OnDelete:CASCADE;"`
	Name     string
	Season   int
	Episode  int
	Duration int
}
