package models

type Movie struct {
	ID          uint   `gorm:"primaryKey"`
	Logo        string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Producer    string `gorm:"not null"`
	Director    string `gorm:"not null"`
	Year        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Seasons     int
	Episodes    int
	Minutes     int
	Screenshots []string    `gorm:"type:text[]"`
	Categories  []*Category `gorm:"many2many:movie_categories;"`
	Genres      []*Genre    `gorm:"many2many:movie_genres"`
	Users       []*User     `gorm:"many2many:user_favorites;"`
}

type MoviesFilter struct {
	Title    string `form:"title"`
	Genre    string `form:"genre"`
	Year     string `form:"year"`
	Type     string `form:"type"`
	SortBy   string `form:"sort_by"`
	SortDesc bool   `form:"sort_desc"`
}

type Video struct {
	ID       uint `gorm:"primarykey"`
	MovieID  uint
	Movie    Movie  `gorm:"constraint:OnDelete:CASCADE;"`
	Name     string `gorm:"not null"`
	URL      string `gorm:"not null"`
	Season   int
	Episode  int
	Duration int `gorm:"not null"`
}
