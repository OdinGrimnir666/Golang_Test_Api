package DataBase

import (
	  "gorm.io/gorm"
  )


type UrlPackage struct {
	gorm.Model
	Id  int
	Url string
}



