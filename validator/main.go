package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type User struct {
	UserName  string   `validate:"required,min=6,max=20" vmsg:"user name must be 6 to 12 bettween"`
	UserPassd string   `validate:"required,min=6,max=12" vmsg:"user passd must be 6 to 12 bettween"`
	TestName  string   `validate:"required,username" vmsg:"user name is wrong"`
	UserTag   []string `validate:"required,min=1,max=4,unique,dive,usertag" vmsg:"user tag illegall"`
}

func main() {
	user := &User{UserName: "122121", UserPassd: "12221212", TestName: "asqwas", UserTag: []string{"aa", "bba", "cca", "dda"}}
	valid := validator.New()
	pattern := "[a-zA-Z]\\w{5,19}"
	err := AddRegexTag("username", pattern, valid)
	if err != nil {
		log.Fatal(err)
	}
	pattern = "^[\u4e00-\u9fa5a-zA-Z0-9]{2,4}$"
	err = AddRegexTag("usertag", pattern, valid)

	err = GetValidMsg(user, valid.Struct(user))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("validate seccess!")
}

func AddRegexTag(tagName string, pattern string, v *validator.Validate) error {
	err := v.RegisterValidation(tagName, func(fl validator.FieldLevel) bool {
		//fmt.Println(fl.Field().String())
		m, err := regexp.MatchString(pattern, fl.Field().String())
		if err != nil {
			log.Fatal(err)
		}

		return m
	}, false)

	return err
}

func GetValidMsg(o interface{}, err error) error {
	getObj := reflect.TypeOf(o)

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
					if value, ok := f.Tag.Lookup("vmsg"); ok {
						return fmt.Errorf("%s", value)
					} else {
						return fmt.Errorf("%s", e)
					}
				} else {
					return fmt.Errorf("%s", e)
				}
			}
		}
	}

	return err
}
