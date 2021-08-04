package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	file.Filename = strconv.Itoa(int(time.Now().Unix())) + "-" + strings.ReplaceAll(file.Filename, " ", "-")
	err = c.SaveUploadedFile(file, "./storage/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusSeeOther, "/file/"+file.Filename)
}

func download(c *gin.Context) {
	c.File("./storage/" + c.Param("filename"))
}

func file(c *gin.Context) {
	c.HTML(http.StatusOK, "file.tmpl", gin.H{"filename": c.Param("filename")})
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("web/*")

	router.GET("/", index)
	router.POST("/upload", upload)

	router.GET("/download/:filename", download)
	router.GET("/file/:filename", file)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":9999")
}
