package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)

func GetBuckets() (*s3.ListBucketsOutput, error) {
	// バケットのリストを取得
	buckets, err := minioClient.ListBuckets(nil)
	if err != nil {
		fmt.Println("Unable to list buckets", err)
		return nil, err
	}

	return buckets, nil
}
