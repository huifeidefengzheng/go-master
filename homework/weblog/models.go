package main

import (
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:50;not null;uniqueIndex" json:"username"`
	Email     string         `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Posts     []Post         `gorm:"foreignKey:UserID" json:"posts"`
	Comments  []Comment      `gorm:"foreignKey:UserID" json:"comments"`
	CreatedAt time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime" json:"deleted_at,omitempty"`
}

// Post 文章模型
type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Comments  []Comment      `gorm:"foreignKey:PostID" json:"comments"`
	CreatedAt time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime" json:"deleted_at,omitempty"`
}

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	PostID    uint           `gorm:"not null;index" json:"post_id"`
	Post      Post           `gorm:"foreignKey:PostID" json:"post"`
	CreatedAt time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime" json:"deleted_at,omitempty"`
}
