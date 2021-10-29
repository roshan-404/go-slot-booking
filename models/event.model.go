package models

type Event struct {
	ID        int `gorm:"primary_key"`
	DateTime string `json:"datetime"`
	Duration int `json:"duration"`
	TimeZone string `json:"timezone"`
	File []File
}

type File struct {
	ID int `gorm:"primary_key"`
	FileURL string `json:"string"`
	EventId int
}
