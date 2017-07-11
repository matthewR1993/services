package database

import (
	"github.com/jinzhu/gorm"
)

var (
	/*
	This global variable provides
	databse connection pool
	*/
	DBCon *gorm.DB
)
