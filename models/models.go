package models

import "gorm.io/gorm"

type User2 struct {
	gorm.Model
	Name           string    `gorm:"size:10"`
	Password       string    `gorm:"size:40"`
	Content        string    `gorm:"size:50"`
	Videos         []Video2   `gorm:"ForeignKey:AuthorID"`
	Comments       []Comment2 `gorm:"many2many:comments;joinForeignKey:UserID"`
	FavoriteVideos []Video2   `gorm:"many2many:user_favorite_videos"`
	Subscribers    []User2    `gorm:"joinForeignKey:SubscriberID;many2many:subscribes"`
	Followers      []User2    `gorm:"joinForeignKey:UserID;many2many:subscribes"`
}

type Video2 struct {
	gorm.Model
	AuthorID      uint
	Title         string    `gorm:"size:30"`
	Author        User2      `gorm:"reference:ID"`
	UserFavorites []User2    `gorm:"many2many:user_favorite_videos"`
	Comments      []Comment2 `gorm:"many2many:Comment;joinForeignKey:VideoID"`
}

type Comment2 struct {
	gorm.Model
	UserID  uint
	VideoID uint
	Content string `gorm:"size:100"`
}

