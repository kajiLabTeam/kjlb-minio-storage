package controller

import (
	"log"
	"net/http"

	"github.com/NenfuAT/24AuthorizationServer/service"
	"github.com/gin-gonic/gin"
)

func GetBuckets(c *gin.Context) {
	type Response struct {
		Buckets []string `json:"buckets"`
	}
	var bucketNames []string

	// サービスからバケットのリストを取得
	buckets, err := service.GetBuckets()
	if err != nil {
		log.Fatalf("Unable to create session, %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to create session",
		})
		return
	}

	// バケットの名前を抽出

	for _, bucket := range buckets.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}
	responseData := Response{
		Buckets: bucketNames,
	}

	// JSONレスポンスを返す
	c.JSON(http.StatusOK, responseData)
}

func CreateBucket(c *gin.Context) {
	bucketName := c.Query("name")

	err := service.CreateBucket(bucketName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"bucket": bucketName,
	})
}