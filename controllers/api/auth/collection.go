package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/trisnaterasoedjita/go-boilerplate-api/models"
)

var (
	db    *gorm.DB
	count int
	user  models.User
)