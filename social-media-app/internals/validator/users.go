package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	
)



func Users(s interface{} )error  {
	err := _validator.Struct(s)
	if err != nil {
		if _,ok :=err.(*validator.ValidationErrors);ok {
			panic(err)
	}
	for _,e := range err.(validator.ValidationErrors) {
		fmt.Println(e.Field())
		fmt.Println(e.Tag())
}
return err
	}
	return nil
}