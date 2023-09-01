package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title         string `gorm:"type:varchar(300)"`
	Subtitle      string `gorm:"type:varchar(1000)"`
	Content       string
	AuthorID      uint
	Author        User
	StatusID      uint
	Status        Status
	Tags          []Tag     `gorm:"many2many:article_tags;"`
	Comments      []Comment `gorm:"many2many:article_comments;"`
	Views         uint
	Likes         uint
	FeaturedImage string
}

type Status struct {
	gorm.Model
	Name     string
	Articles []Article ` gorm:"many2many:article_status;"`
}

type Tag struct {
	gorm.Model
	Name     string
	Articles []Article `gorm:"many2many:article_tags;"`
}

type Comment struct {
	gorm.Model
	ArticleID uint
	Content   string
	AuthorID  uint
	Author    User
}
