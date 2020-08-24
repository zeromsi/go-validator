package v1

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
)

const(
	TRUE="true"
	FALSE="false"
	TYPE="type"
	EMAIL="email"
	STRING="string"
	REQUIRED="required"
	MESSAGE="msg"
)

type Validator struct {
	l * log.Logger
}

func NewValidatorWithLogger (l *log.Logger) *Validator{
	return &Validator{l}
}

func NewValidator () *Validator{
	return &Validator{}
}


func(v Validator) Struct(object interface{}) [] string{
	if v.l!=nil{
		log.Println(v.l.Prefix())
	}

	errs:=[] string{}
	val := reflect.ValueOf(object)
	for i := 0; i < val.Type().NumField(); i++ {
		required:=val.Type().Field(i).Tag.Get(REQUIRED)
		dataType:=val.Type().Field(i).Tag.Get(TYPE)
		value := fmt.Sprintf("%v",val.Field(i))
		msg:=val.Type().Field(i).Tag.Get(MESSAGE)
		if(required=="true"){
			if msg==""{

				if msg==""{
					errs = append(errs,"Please provide a name![ERROR]: empty name!" )
				}else{
					errs = append(errs, "[ERROR]:"+msg)
				}
				if dataType==EMAIL{
					if v.checkIfEmailIsValid(value)==false{
						if msg==""{
							errs = append(errs, "Please provide a valid email! [ERROR]:Invalid email "+value)
						}else{
							errs = append(errs, "[ERROR]:"+msg)
						}

					}
				}

			}else{
				if dataType==EMAIL{
					if v.checkIfEmailIsValid(value)==false{
						if msg==""{
							errs = append(errs, "Please provide a valid email! [ERROR]:Invalid email "+value)
						}else{
							errs = append(errs, "[ERROR]:"+msg)
						}

					}
				}
			}


		}

	}


return errs

}

func (v Validator) checkIfEmailIsValid(email string) bool{
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}
