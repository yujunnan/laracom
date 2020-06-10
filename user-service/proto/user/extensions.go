package laracom_service_user

import (
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//BeforeCreate
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	fmt.Println(uuid.String())
	return scope.SetColumn("Id", uuid.String())
}
