package models

type Movie struct {
	ID          uint     `gorm:"primarykey"`
	Logo        string   `gorm:"not null"`
	Title       string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	Producer    string   `gorm:"not null"`
	Director    string   `gorm:"not null"`
	Year        string   `gorm:"not null"`
	Type        string   `gorm:"not null"`
	Tags        []string `gorm:"type:text[]"`
	Seasons     int      `gorm:"not null"`
	Episodes    int      `gorm:"not null"`
	Minutes     int
	Screenshots []string `gorm:"type:text[]"`
	Users       []*User  `gorm:"many2many:user_favorites;"`
}

type MoviesFilter struct {
	
}

type Video struct {
	ID       uint `gorm:"primarykey"`
	MovieID  uint
	Movie    Movie  `gorm:"constraint:OnDelete:CASCADE;"`
	Name     string `gorm:"not null"`
	Ref      string `gorm:"not null"`
	Season   int
	Episode  int
	Duration int `gorm:"not null"`
}
