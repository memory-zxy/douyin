package dao

import "time"

type Video struct {
	VideoID       uint64    `gorm:"column:video_id;primary_key;NOT NULL" redis:"-"`
	Title         string    `gorm:"column:title;NOT NULL" redis:"title"`
	PlayName      string    `gorm:"column:play_name;NOT NULL" redis:"play_name"`
	CoverName     string    `gorm:"column:cover_name;NOT NULL" redis:"cover_name"`
	FavoriteCount int64     `gorm:"-" redis:"favorite_count"`
	CommentCount  int64     `gorm:"-" redis:"comment_count"`
	AuthorID      uint64    `gorm:"column:author_id;NOT NULL" redis:"author_id"`
	CreatedAt     time.Time `gorm:"column:created_at" redis:"-"`
	ExtInfo       *string   `gorm:"column:ext_info" redis:"-"`
}

type VideoID struct {
	VideoID uint64
}

func (VideoID) TableName() string {
	return "videos"
}
