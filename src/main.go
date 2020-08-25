package main

import (
	v1 "github.com/zeromsi/go.validator/src/v1"
	"log"
)

func main(){
emp:=Employee{
	Email: "abc"}.Validate()
log.Println(emp)
}

type Employee struct {
	Name  string `required:"true" msg:"Please provide a name"`
	Email string `required:"true" type:"email"`
}

func (employee Employee) Validate() []string {
	//l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	errs := v1.NewValidator().Struct(employee)
	return errs
}