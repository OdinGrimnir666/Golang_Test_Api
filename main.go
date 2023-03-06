package main

import (
	"fmt"
	"package/apirequest"
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Starting..")

	router := gin.New()
	router.SetTrustedProxies(nil)

	router.GET("/",apirequest.GetUrlPackage)

	router.POST("/",apirequest.PostUrlPackage)

    router.Run("localhost:8080")

}
