package table

import (
	"encoding/json"
	"fmt"
	allconst "info-end/const"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
)

type VideoRecord struct {
	Rid      string   `json:"rid"`
	UserName string   `json:"user_name"`
	Vname    string   `json:"vname"`
	Vdesc    string   `json:"vdesc"`
	Vurl     string   `json:"vurl"`
	Vid      string   `json:"vid"`
	Vmurl    string   `json:"vmurl"`
	Vmid     string   `json:"vmid"`
	Vtag     []string `json:"vtag"`
	Vtime    int64    `json:"vtime"`
	Vstatus  string   `json:"vstatus"`
}

//! 创建视频记录
func CreateVideoRecord(client *tablestore.TableStoreClient, data VideoRecord) (interface{}, error) {
	rand.Seed(time.Now().Unix())
	newData := VideoRecord{
		Rid:      fmt.Sprint(rand.Int63()),
		UserName: data.UserName,
		Vname:    data.Vname,
		Vdesc:    data.Vdesc,
		Vurl:     data.Vurl,
		Vid:      data.Vid,
		Vmurl:    data.Vmurl,
		Vmid:     data.Vmid,
		Vtag:     data.Vtag,
		Vtime:    data.Vtime,
		Vstatus:  data.Vstatus,
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = allconst.Tables["video_record"]
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("rid", newData.Rid)
	putPk.AddPrimaryKeyColumn("username", newData.UserName)
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("vname", newData.Vname)
	putRowChange.AddColumn("vdesc", newData.Vdesc)
	putRowChange.AddColumn("vurl", newData.Vurl)
	putRowChange.AddColumn("vid", newData.Vid)
	putRowChange.AddColumn("vmurl", newData.Vmurl)
	putRowChange.AddColumn("vmid", newData.Vmid)
	putRowChange.AddColumn("vtime", newData.Vtime)
	putRowChange.AddColumn("vvstatus", newData.Vstatus)
	arr, _ := json.Marshal(newData.Vtag)
	putRowChange.AddColumn("vtag", string(arr))
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("CreateVideoRecord:", err)
		return data, err
	} else {
		fmt.Println("CreateVideoRecord finished", res)
	}
	// newRes, _ := GetAllVideoRecord(client)
	return newData, nil
}

type SearchParams struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
}

//! 列表查询记录
func GetAllVideoRecord(client *tablestore.TableStoreClient, searchBody SearchParams, userType interface{}, username interface{}) (interface{}, error) {
	page, _ := strconv.ParseInt(searchBody.Page, 10, 32)
	pageSize, _ := strconv.ParseInt(searchBody.PageSize, 10, 32)
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(allconst.Tables["video_record"])
	searchRequest.SetIndexName(allconst.Table_index["video_record"])
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
		res := handleGetAllVideoRecordRes(searchResponse)
		if err != nil {
			fmt.Println("GetAllVideoRecord batchget failed with error:", err)
			return nil, err
		} else {
			fmt.Println("GetAllVideoRecord batchget finished")
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
		res := handleGetAllVideoRecordRes(searchResponse)
		if err != nil {
			fmt.Println("GetAllVideoRecord batchget failed with error:", err)
			return nil, err
		} else {
			fmt.Println("GetAllVideoRecord batchget finished")
		}
		return res, nil
	}
}

type VideoRecordRes struct {
	Data  []VideoRecord `json:"data"`
	Count int64         `json:"count"`
}

func handleGetAllVideoRecordRes(originData *tablestore.SearchResponse) VideoRecordRes {
	count := originData.TotalCount
	var res []VideoRecord
	for _, item := range originData.Rows {
		var temp VideoRecord
		json.Unmarshal([]byte(item.Columns[5].Value.(string)), &temp.Vtag)
		temp.Rid = item.PrimaryKey.PrimaryKeys[0].Value.(string)
		temp.UserName = item.PrimaryKey.PrimaryKeys[1].Value.(string)
		temp.Vname = item.Columns[4].Value.(string)
		temp.Vdesc = item.Columns[0].Value.(string)
		temp.Vurl = item.Columns[7].Value.(string)
		temp.Vid = item.Columns[1].Value.(string)
		temp.Vmurl = item.Columns[3].Value.(string)
		temp.Vmid = item.Columns[2].Value.(string)
		if len(item.Columns) > 8 {
			temp.Vstatus = item.Columns[8].Value.(string)
		} else {
			temp.Vstatus = "init"
		}
		temp.Vtime = item.Columns[6].Value.(int64)
		res = append(res, temp)
	}
	return VideoRecordRes{
		Data:  res,
		Count: count,
	}
}

//! 根据rid查询视频记录
func SearchVideorecordByRid(client *tablestore.TableStoreClient, rid string) (interface{}, error) {
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(allconst.Tables["video_record"])
	searchRequest.SetIndexName(allconst.Table_index["video_record"])
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
		fmt.Println("GetAllVideoRecord batchget failed with error:", err)
		return nil, err
	} else {
		fmt.Println("GetAllVideoRecord batchget finished")
	}
	res := handleGetAllVideoRecordRes(searchResponse)
	return res.Data[0], nil
}

//! 根据id更新视频记录
func UpdateVideoRecord(client *tablestore.TableStoreClient, data VideoRecord) (interface{}, error) {
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = allconst.Tables["video_record"]
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("rid", data.Rid)
	updatePk.AddPrimaryKeyColumn("username", data.UserName)
	updateRowChange.PrimaryKey = updatePk
	updateRowChange.PutColumn("vname", data.Vname)
	updateRowChange.PutColumn("vdesc", data.Vdesc)
	updateRowChange.PutColumn("vurl", data.Vurl)
	updateRowChange.PutColumn("vid", data.Vid)
	updateRowChange.PutColumn("vmurl", data.Vmurl)
	updateRowChange.PutColumn("vmid", data.Vmid)
	updateRowChange.PutColumn("vtime", data.Vtime)
	updateRowChange.PutColumn("vvstatus", data.Vstatus)
	arr, _ := json.Marshal(data.Vtag)
	updateRowChange.PutColumn("vtag", string(arr))
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	updateRowRequest.UpdateRowChange = updateRowChange
	res, err := client.UpdateRow(updateRowRequest)

	if err != nil {
		fmt.Println("upDateVideoRecord:", err)
		return res, err
	} else {
		fmt.Println("update row finished")
	}
	return res, nil
}

//! 删除一条视频记录
func DelVideoRecord(client *tablestore.TableStoreClient, rid string, username interface{}) (interface{}, error) {
	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = allconst.Tables["video_record"]
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("rid", rid)
	deletePk.AddPrimaryKeyColumn("username", username)
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)
	_, err := client.DeleteRow(deleteRowReq)
	if err != nil {
		fmt.Println("DelVideoRecord error:", err)
		return nil, err
	} else {
		fmt.Println("DelVideoRecord finished")
	}
	return nil, nil
}
