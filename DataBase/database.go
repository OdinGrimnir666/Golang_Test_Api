package DataBase

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func DB() (gorm.DB){
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
		panic("failed to connect database")
	}
  return *db
}

func RequestGet(id int)(UrlPackage){
  var urlpackage UrlPackage
  var db gorm.DB = DB()
  db.First(&urlpackage ,id)
  return urlpackage
}


func RequestPostArray(arrayurlpackage [] int)([]UrlPackage){
  var urlpackage []UrlPackage
  var db gorm.DB = DB()
  db.Find(&urlpackage, arrayurlpackage)
  return urlpackage
}



func DateInit() {
  var db gorm.DB = DB()
  db.AutoMigrate(&UrlPackage{})
  db.Create(&UrlPackage{Id: 1,Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1"})
  db.Create(&UrlPackage{Id: 2,Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=2"})
  db.Create(&UrlPackage{Id: 3,Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=3"})
  db.Create(&UrlPackage{Id: 4,Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=4"})
}









