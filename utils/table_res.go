package utils

import (
	"fmt"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func HandleTableRes(data *tablestore.GetRowResponse) map[string]interface{} {
	PrimaryKey := data.PrimaryKey.PrimaryKeys
	Columns := data.Columns
	res := make(map[string]interface{})
	for _, v := range PrimaryKey {
		res[v.ColumnName] = v.Value
	}
	for _, v := range Columns {
		res[v.ColumnName] = v.Value
	}
	return res
}

type AvatorItem struct {
	Uid        interface{} `json:"uid"`
	Avator_url interface{} `json:"avator_url"`
}

func HandleTableResBatch(data *tablestore.GetRangeResponse) *[]AvatorItem {
	res := make([]AvatorItem, 0)
	for _, v := range data.Rows {
		data := AvatorItem{
			Uid:        fmt.Sprint(v.PrimaryKey.PrimaryKeys[0].Value),
			Avator_url: fmt.Sprint(v.Columns[0].Value),
		}
		res = append(res, data)
	}

	return &res
}
