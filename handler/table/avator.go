package table

import (
	"fmt"
	allconst "info-end/const"
	"info-end/utils"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func AvatorGet(client *tablestore.TableStoreClient) (interface{}, error) {
	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	rangeRowQueryCriteria.TableName = allconst.Tables["avator"]

	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumnWithMinValue("avator_id")
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumnWithMaxValue("avator_id")
	rangeRowQueryCriteria.StartPrimaryKey = startPK
	rangeRowQueryCriteria.EndPrimaryKey = endPK
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	rangeRowQueryCriteria.MaxVersion = 1
	rangeRowQueryCriteria.Limit = 99999
	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria

	getRangeResp, err := client.GetRange(getRangeRequest)

	res := utils.HandleTableResBatch(getRangeResp)
	if err != nil {
		fmt.Println("batchget failed with error:", err)
		return nil, err
	} else {
		fmt.Println("batchget finished")
		return res, nil
	}
}
