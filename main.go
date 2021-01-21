package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		dst := "/var/www/html/repos/" + files[0].Filename
		c.SaveUploadedFile(files[0], dst)

		dst = "/var/www/html/data/version/" + files[1].Filename
		c.SaveUploadedFile(files[1], dst)

		dst = "/var/www/html/data/depend/" + files[2].Filename
		c.SaveUploadedFile(files[2], dst)

		c.String(http.StatusOK, fmt.Sprintf("File %s and %s uploaded!", files[0].Filename, files[1].Filename, files[2].Filename))
	})

	r.GET("/pass", func(c *gin.Context) {
		haslo := c.DefaultQuery("haslo", "")
		packageName := c.Query("name")

		c.String(http.StatusOK, "%s", packageName)

		f, _ := os.Create("/home/pibu/vanilla-server/passwords/" + string(packageName) + ".txt")
		_, _ = f.WriteString(string(haslo))
	})

	r.Run(":2137")
}
