package example

import (
	v1 "github.com/zeromsi/go.validator/src/v1"
	"log"
	"os"
)

type Employee struct {
	Name  string `required:"true" msg:"Please provide a name" `
	Email string `required:"true" type:"email"`
}

func (employee Employee) Validate() []string {
	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	errs := v1.NewValidatorWithLogger(l).Struct(employee)
	return errs
}
