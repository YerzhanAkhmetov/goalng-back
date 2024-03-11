package models

import "gorm.io/gorm"

//Здесь объекты в которых потом будет лежать данные
type Post struct {
	gorm.Model
	Title     string
	Body      string
	ImagePath string
}
type User struct {
	gorm.Model
	Name        string
	Surname     string
	Age         int
	Email       string
	NumberPhone int
}
