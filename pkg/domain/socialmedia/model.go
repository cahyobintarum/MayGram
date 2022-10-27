package socialmedia

import (
	"github.com/cahyobintarum/MayGram/pkg/domain/gormmodel"
)

type SocialMedia struct {
	gormmodel.GormModel
	Name   string `gorm:"not null" json:"name" valid:"required~social media name is required"`
	URL    string `gorm:"not null" json:"url" valid:"required~social media url is required,url~invalid url format"`
	UserID uint64 `json:"user_id"`
}
