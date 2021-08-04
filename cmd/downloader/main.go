package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func download(c *gin.Context){
	c.File("./storage/"+c.Param("filename"))
}

func file(c *gin.Context){
	c.HTML(http.StatusOK, "file.tmpl", gin.H{"filename":c.Param("filename")})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.GET("/download/:filename", download)
	router.GET("/file/:filename", file)
	router.Run(":9999")
}
