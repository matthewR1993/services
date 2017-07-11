package valid

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	/*
	This global variable provides
	validator object with caching
	so only one instance preferable
	*/
	Validate *validator.Validate
)
