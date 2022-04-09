package table

import (
	"fmt"
	allconst "info-end/const"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
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

type TopicRecord struct {
	Rid       string `json:"rid"`
	UserName  string `json:"user_name"`
	TopicType string `json:"topic_type"`
	TopicBody string `json:"topic_body"`
}

func CreateTopicRecord(client *tablestore.TableStoreClient, data TopicRecord) (interface{}, error) {
	rand.Seed(time.Now().Unix())
	newData := TopicRecord{
		Rid:       fmt.Sprint(rand.Int63()),
		UserName:  data.UserName,
		TopicType: data.TopicType,
		TopicBody: data.TopicBody,
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = allconst.Tables["topic_record"]
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("rid", newData.Rid)
	putPk.AddPrimaryKeyColumn("username", newData.UserName)
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("topic_type", newData.TopicType)
	putRowChange.AddColumn("topic_body", newData.TopicBody)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)
	if err != nil {
		fmt.Println("CreateVideoRecord:", err)
		return data, err
	} else {
		fmt.Println("CreateVideoRecord finished", res)
	}
	return newData, nil
}
