package validator

import (
	"github.com/go-playground/validator/v10"
	

)

var _validator *validator.Validate

func init(){
	_validator = validator.New()
}
func Validate(s interface{}) error {
	return _validator.Struct(s)

}
type Payload struct {
    Field1 string `validate:"required"`
    Field2 int  `validate:"gte=0"`  
}
