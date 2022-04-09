package table

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	allconst "info-end/const"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func UploadUserAvator(filename string, filepath string) (string, error) {
	log.Println("OSS Go SDK Version: ", oss.Version)
	client, err := oss.New(allconst.OssEndpoint, allconst.AccessKeyId, allconst.AccessKeySecret)
	if err != nil {
		fmt.Println("Init oss error:", err)
		return "", err
	}
	bucket, _ := client.Bucket(allconst.UserAvatorBucketName)
	fmt.Println("filepath", filepath)
	err = bucket.PutObjectFromFile("avator/"+filename, filepath)
	if err != nil {
		fmt.Println("upload oss error:", err)
		return "", err
	}
	err = os.Remove(filepath)
	if err != nil {
		fmt.Println("os.Remove error:", err)
		return "", err
	}
	resPath := allconst.AvatorPrefix + "/" + filename
	return resPath, nil
}

//! 管理员管理头像 添加
func AvatorAdd(client *tablestore.TableStoreClient, resPath string) (interface{}, error) {
	rand.Seed(time.Now().Unix())
	uid := rand.Int63()
	data := map[string]interface{}{
		"uid":    uid,
		"avator": resPath,
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = allconst.Tables["avator"]
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("avator_id", fmt.Sprint(data["uid"]))
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("avayor_url", data["avator"])
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)
	if err != nil {
		fmt.Println("AvatorAdd error:", err)
		return data, err
	} else {
		fmt.Println("AvatorAdd finished", res)
	}
	return data, nil
}

//! 管理员管理头像 删除

func AvatorDel(client *tablestore.TableStoreClient, uid string) error {
	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = allconst.Tables["avator"]
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("avator_id", uid)
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)
	_, err := client.DeleteRow(deleteRowReq)
	if err != nil {
		fmt.Println("AvatorDel error:", err)
		return err
	} else {
		fmt.Println("AvatorDel finished")
	}
	return nil
}

func UploadUserVideo(filename string, filepath string) (string, error) {
	log.Println("OSS Go SDK Version: ", oss.Version)
	client, err := oss.New(allconst.OssEndpoint, allconst.AccessKeyId, allconst.AccessKeySecret)
	if err != nil {
		fmt.Println("Init oss error:", err)
		return "", err
	}
	bucket, _ := client.Bucket(allconst.UserAvatorBucketName)
	fmt.Println("filepath", filepath)
	err = bucket.PutObjectFromFile("video/"+filename, filepath)
	if err != nil {
		fmt.Println("upload oss error:", err)
		return "", err
	}
	err = os.Remove(filepath)
	if err != nil {
		fmt.Println("os.Remove error:", err)
		return "", err
	}
	resPath := allconst.VideoPrefix + "/" + filename
	return resPath, nil
}

//! 上传视频并存储视频
func VideoUpload(client *tablestore.TableStoreClient, resPath string) (interface{}, error) {
	rand.Seed(time.Now().Unix())
	uid := rand.Int63()
	data := map[string]interface{}{
		"vid":  uid,
		"vurl": resPath,
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = allconst.Tables["video"]
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("vid", fmt.Sprint(data["vid"]))
	putPk.AddPrimaryKeyColumn("vurl", data["vurl"])
	putRowChange.PrimaryKey = putPk
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)
	if err != nil {
		fmt.Println("VideoUpload error:", err)
		return data, err
	} else {
		fmt.Println("VideoUpload finished", res)
	}
	return data, nil
}

//! 上传视频封面并存储
func VideoImgUpload(client *tablestore.TableStoreClient, resPath string) (interface{}, error) {
	rand.Seed(time.Now().Unix())
	uid := rand.Int63()
	data := map[string]interface{}{
		"vmid":  uid,
		"vmurl": resPath,
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = allconst.Tables["video_img"]
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("vmid", fmt.Sprint(data["vmid"]))
	putPk.AddPrimaryKeyColumn("vmurl", data["vmurl"])
	putRowChange.PrimaryKey = putPk
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)
	if err != nil {
		fmt.Println("VideoUpload error:", err)
		return data, err
	} else {
		fmt.Println("VideoUpload finished", res)
	}
	return data, nil
}
