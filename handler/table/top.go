package table

import (
	"fmt"
	allconst "info-end/const"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
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

type TopicSearchParams struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
}

func GetAllTopicRecord(client *tablestore.TableStoreClient, searchBody TopicSearchParams, userType interface{}, username interface{}) (interface{}, error) {
	page, _ := strconv.ParseInt(searchBody.Page, 10, 32)
	pageSize, _ := strconv.ParseInt(searchBody.PageSize, 10, 32)
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(allconst.Tables["topic_record"])
	searchRequest.SetIndexName(allconst.Table_index["topic_record"])
	fmt.Println(userType, username)
	if userType == "admin" {
		query := &search.MatchAllQuery{}
		searchQuery := search.NewSearchQuery()
		searchQuery.SetQuery(query)
		searchQuery.SetGetTotalCount(true)
		searchQuery.SetLimit(int32(pageSize))
		searchQuery.SetOffset(int32(page-1) * int32(pageSize))
		searchRequest.SetSearchQuery(searchQuery)
		searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
			ReturnAll: true,
		})
		searchResponse, err := client.Search(searchRequest)
		res := handleGetAllTopicRecordRes(searchResponse)
		if err != nil {
			fmt.Println("GetAllTopicRecord batchget failed with error:", err)
			return nil, err
		} else {
			fmt.Println("GetAllTopicRecord batchget finished")
		}
		return res, nil
	} else {
		query := &search.TermQuery{}
		query.FieldName = "username"
		query.Term = username
		searchQuery := search.NewSearchQuery()
		searchQuery.SetQuery(query)
		searchQuery.SetGetTotalCount(true)
		searchQuery.SetLimit(int32(pageSize))
		searchQuery.SetOffset(int32(page-1) * int32(pageSize))
		searchRequest.SetSearchQuery(searchQuery)
		searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
			ReturnAll: true,
		})
		searchResponse, err := client.Search(searchRequest)
		res := handleGetAllTopicRecordRes(searchResponse)
		if err != nil {
			fmt.Println("GetAllTopicRecord batchget failed with error:", err)
			return nil, err
		} else {
			fmt.Println("GetAllTopicRecord batchget finished")
		}
		return res, nil
	}
}

type TopicRecordRes struct {
	Data  []TopicRecord `json:"data"`
	Count int64         `json:"count"`
}

func handleGetAllTopicRecordRes(originData *tablestore.SearchResponse) TopicRecordRes {
	count := originData.TotalCount
	var res []TopicRecord
	for _, item := range originData.Rows {
		var temp TopicRecord
		temp.Rid = item.PrimaryKey.PrimaryKeys[0].Value.(string)
		temp.UserName = item.PrimaryKey.PrimaryKeys[1].Value.(string)
		temp.TopicType = item.Columns[1].Value.(string)
		temp.TopicBody = item.Columns[0].Value.(string)
		res = append(res, temp)
	}
	return TopicRecordRes{
		Data:  res,
		Count: count,
	}
}

func SearchTopicRecordByRid(client *tablestore.TableStoreClient, rid string) (interface{}, error) {
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(allconst.Tables["topic_record"])
	searchRequest.SetIndexName(allconst.Table_index["topic_record"])
	query := &search.TermQuery{}
	query.FieldName = "rid"
	query.Term = rid
	searchQuery := search.NewSearchQuery()
	searchQuery.SetQuery(query)
	searchQuery.SetGetTotalCount(true)
	searchRequest.SetSearchQuery(searchQuery)
	searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
		ReturnAll: true,
	})
	searchResponse, err := client.Search(searchRequest)
	if err != nil {
		fmt.Println("SearchTopicRecordByRid batchget failed with error:", err)
		return nil, err
	} else {
		fmt.Println("SearchTopicRecordByRid batchget finished")
	}
	res := handleGetAllTopicRecordRes(searchResponse)
	return res.Data[0], nil
}

func UpdateTopicRecord(client *tablestore.TableStoreClient, data TopicRecord) (interface{}, error) {
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = allconst.Tables["topic_record"]
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("rid", data.Rid)
	updatePk.AddPrimaryKeyColumn("username", data.UserName)
	updateRowChange.PrimaryKey = updatePk
	updateRowChange.PutColumn("topic_type", data.TopicType)
	updateRowChange.PutColumn("topic_body", data.TopicBody)
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	updateRowRequest.UpdateRowChange = updateRowChange
	res, err := client.UpdateRow(updateRowRequest)

	if err != nil {
		fmt.Println("UpdateTopicRecord:", err)
		return res, err
	} else {
		fmt.Println("update row finished")
	}
	return res, nil
}

func DelTopicRecord(client *tablestore.TableStoreClient, rid string, username interface{}) (interface{}, error) {
	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = allconst.Tables["topic_record"]
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("rid", rid)
	deletePk.AddPrimaryKeyColumn("username", username)
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)
	_, err := client.DeleteRow(deleteRowReq)
	if err != nil {
		fmt.Println("DelTopicRecord error:", err)
		return nil, err
	} else {
		fmt.Println("DelTopicRecord finished")
	}
	return nil, nil
}
