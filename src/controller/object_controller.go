package controller

import (
	"net/http"

	"github.com/NenfuAT/24AuthorizationServer/helper"
	"github.com/NenfuAT/24AuthorizationServer/service"
	"github.com/gin-gonic/gin"
)

func PostObject(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	err := helper.AuthBasic(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	bucketName := c.Request.FormValue("bucket")
	if bucketName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bucket name is required",
		})
		return
	}

	path := c.Request.FormValue("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "path name is required",
		})
		return
	}

	// フォームからファイルを取得
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to get file",
		})
		return
	}
	defer file.Close()
	fileName := header.Filename

	err = service.PostObject(bucketName, path, file, fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"bucket": bucketName,
		"path":   path,
		"file":   fileName,
	})
}

func GetObjectUrl(c *gin.Context) {
	type Request struct {
		Bucket string `json:"bucket"`
		Path   string `json:"path"`
		File   string `json:"file"`
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	err := helper.AuthBasic(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req Request
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := service.GetObjectUrl(req.Bucket, req.Path, req.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
