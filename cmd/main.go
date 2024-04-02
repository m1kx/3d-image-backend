package main

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m1kx/image/internal/gif"
	"github.com/m1kx/image/util"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	})
	r.POST("/api/gif", func(c *gin.Context) {
		var points []util.Vec2
		points_str, exists := c.GetPostForm("points")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Please provide points",
			})
			return
		}
		err := json.Unmarshal([]byte(points_str), &points)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err,
			})
			return
		}
		formfile, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		file, err := formfile.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		imageInput, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		gifBytes, err := gif.CreateGif(&imageInput, &points)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		c.Writer.Write(gifBytes)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.Run(":9999")
}
