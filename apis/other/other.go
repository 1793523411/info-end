package other

import (
	"encoding/json"
	"fmt"
	allconst "info-end/const"
	"math/rand"
	"time"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
	"github.com/gin-gonic/gin"
)

func SearchDemo(c *gin.Context) {
	res, _ := addData()
	res = TermQuery(allconst.Client, "test2", "test2_index")
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
	})
}

/**
 * 查询表中Col_Keyword列精确匹配"hangzhou"的数据。
 */
func TermQuery(client *tablestore.TableStoreClient, tableName string, indexName string) interface{} {
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(tableName)
	searchRequest.SetIndexName(indexName)
	query := &search.MatchAllQuery{} //设置查询类型为MatchAllQuery。
	searchQuery := search.NewSearchQuery()
	searchQuery.SetQuery(query)
	searchQuery.SetGetTotalCount(true)
	// searchQuery.SetLimit(1) //设置Limit为0，表示不获取具体数据。
	searchQuery.SetLimit(10)
	searchQuery.SetOffset(5)
	searchRequest.SetSearchQuery(searchQuery)
	//设置为返回所有列。
	searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
		ReturnAll: true,
	})
	searchResponse, err := client.Search(searchRequest)
	if err != nil {
		fmt.Printf("%#v", err)
		return nil
	}
	fmt.Println("IsAllSuccess: ", searchResponse.IsAllSuccess) //查看返回结果是否完整。
	fmt.Println("TotalCount: ", searchResponse.TotalCount)     //打印匹配到的总行数，非返回行数。
	fmt.Println("RowCount: ", len(searchResponse.Rows))
	for _, row := range searchResponse.Rows {
		jsonBody, err := json.Marshal(row)
		if err != nil {
			panic(err)
		}
		fmt.Println("Row: ", string(jsonBody))
	}
	return searchResponse.Rows
}

func addData() (interface{}, error) {
	rand.Seed(time.Now().Unix())
	client := allconst.Client
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "test2"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("uid", "111"+fmt.Sprint(rand.Int63()))
	putPk.AddPrimaryKeyColumn("name", "testpk1"+fmt.Sprint(rand.Int63()))
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("username", "111"+fmt.Sprint(rand.Int63()))
	putRowChange.AddColumn("password", "2121"+fmt.Sprint(rand.Int63()))
	putRowChange.AddColumn("type", "212"+fmt.Sprint(rand.Int63()))
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("addData:", err)
		return res, err
	} else {
		fmt.Println("addData finished", res)
	}
	return res, nil
}
