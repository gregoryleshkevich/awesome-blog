package services

import (
  	"github.com/jinzhu/gorm"
	"awesome-blog/internal/models"
)

type DB struct {
	Connect *gorm.DB
}

func (db *DB) create(title, body string) *models.Post {
	post := &models.Post{Title: title, Body: body}
	db.Connect.Create(post)
	return post
}

