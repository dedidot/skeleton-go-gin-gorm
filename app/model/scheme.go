package model
import (
	"github.com/jinzhu/gorm"
)

type Scheme struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Scheme) TableName() string {
	return "TABLE_NAME"
}