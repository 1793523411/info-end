package table

import (
	"fmt"
	allconst "info-end/const"
	"log"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadTopImg(filename string, filepath string) (string, error) {
	log.Println("OSS Go SDK Version: ", oss.Version)
	client, err := oss.New(allconst.OssEndpoint, allconst.AccessKeyId, allconst.AccessKeySecret)
	if err != nil {
		fmt.Println("Init oss error:", err)
		return "", err
	}
	bucket, _ := client.Bucket(allconst.UserAvatorBucketName)
	fmt.Println("filepath", filepath)
	err = bucket.PutObjectFromFile("topImg/"+filename, filepath)
	if err != nil {
		fmt.Println("upload oss error:", err)
		return "", err
	}
	err = os.Remove(filepath)
	if err != nil {
		fmt.Println("os.Remove error:", err)
		return "", err
	}
	resPath := allconst.TopUrlPrefix + "/" + filename
	return resPath, nil
}
