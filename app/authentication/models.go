package authentication

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}
