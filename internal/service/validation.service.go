package service

import "github.com/go-playground/validator/v10"

var Validator validationStruct

type validationStruct struct{}

func (v *validationStruct) validate(a any) error {
	return validator.New().Struct(a)
}
