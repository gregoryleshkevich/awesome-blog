package models

import (
	"github.com/jinzhu/gorm"
)

// Post model
type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PostDBContract
type PostDBContract interface {
	GetByID(ID int) *Post
	GetAll() []*Post
	Create(post *Post) *Post
	Update(post *Post) *Post
	Delete(post *Post)
}

// PostDB struct
type PostDB struct{}

// GetByID returns post by ID
func (pd PostDB) GetByID(ID int) *Post {
	var post Post
	db.First(&post, ID)
	return &post
}

// GetAll posts
func (pd PostDB) GetAll() []*Post {
	var posts []*Post
	db.Find(&posts)

	return posts
}

// Create post
func (pd PostDB) Create(post *Post) *Post {
	db.Create(post)

	return post
}

// Update post
func (pd PostDB) Update(post *Post) *Post {
	db.Save(post)

	return post
}

// Delete post
func (pd PostDB) Delete(post *Post) {
	db.Delete(post)
}