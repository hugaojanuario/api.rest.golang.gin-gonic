package models

import (

"gorm.io/gorm"
"gopkg.in/validator.v2"
)

type Student struct {
	gorm.Model

	Name string `json:"name" validate:"nonzero, regexp=^[a-zA-Z]*$"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}


func ValidationStudent(s *Student) error{
	if err := validator.Validate(s); err !=nil{
		return err
	}
	return nil
}