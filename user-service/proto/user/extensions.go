package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	fmt.Println(uuid.String())
	return scope.SetColumn("Id", uuid.String())
}