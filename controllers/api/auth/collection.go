package auth

import (
	"github.com/jinzhu/gorm"
	"go-boilerplate-api/models"
)

var (
	db    *gorm.DB
	count int
	user  models.User
)