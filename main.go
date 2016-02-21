package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"

	"github.com/Depado/portfolio/utils"
)

func main() {
	var err error

	tbox, _ := rice.FindBox("templates")
	abox, _ := rice.FindBox("assets")

	r := gin.Default()
	if err = utils.InitAssetsTemplates(r, tbox, abox, "index.html"); err != nil {
		log.Fatal(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run(":8093")
}
