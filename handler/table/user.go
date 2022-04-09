package table

import (
	"errors"
	"fmt"
	allconst "info-end/const"
	"info-end/middleware"
	"info-end/utils"
	"math/rand"
	"time"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
	"github.com/golang-jwt/jwt"
)

//! 创建数据表
func CreateTableSample(client *tablestore.TableStoreClient, tableName string) {
	createTableRequest := new(tablestore.CreateTableRequest)
	//创建主键列的schema，包括PK的个数、名称和类型。
	//第一个PK列为整数，名称是pk0，此列同时也是分区键。
	//第二个PK列为整数，名称是pk1。
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk0", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createTableRequest.TableMeta = tableMeta
	createTableRequest.TableOption = tableOption
	createTableRequest.ReservedThroughput = reservedThroughput
	_, err := client.CreateTable(createTableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}
}

type User struct {
	Uid      int64  `json:"u_id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Type     string `json:"type"`
}

//! 注册用户
func InsertUser(client *tablestore.TableStoreClient, tableName string, data User) (User, error) {
	rand.Seed(time.Now().Unix())
	isExit, err := SearchUseamerIsExit(client, data.UserName)
	data = User{
		Uid:      rand.Int63(),
		UserName: data.UserName,
		PassWord: data.PassWord,
		Type:     "common",
	}
	if err != nil {
		return data, err
	}
	if isExit {
		return data, errors.New("用户已存在")
	}
	rand.Seed(time.Now().Unix())
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = tableName
	putPk := new(tablestore.PrimaryKey)

	putPk.AddPrimaryKeyColumn("username", data.UserName)
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("uid", data.Uid)
	putRowChange.AddColumn("password", data.PassWord)
	putRowChange.AddColumn("type", data.Type)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("InsertUser:", err)
		return data, err
	} else {
		fmt.Println("InsertUser finished", res)
	}
	return data, nil
}

//!查找用户是否存在
func SearchUseamerIsExit(client *tablestore.TableStoreClient, username string) (bool, error) {
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("username", username)

	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = "user"
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, err := client.GetRow(getRowRequest)

	if err != nil {
		fmt.Println("SearchUseamerIsExit  error:", err)
		return false, err
	} else {
		return len(getResp.Columns) != 0, nil
	}
}

//! 用户登录逻辑
type loginInfo struct {
	Uid      int64  `json:"u_id"`
	UserName string `json:"username"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Token    string `json:"token"`
}

func Login(client *tablestore.TableStoreClient, username string, password string) (loginInfo, error) {
	data := loginInfo{
		Uid:      0,
		UserName: "",
		Status:   "fail",
		Type:     "",
	}
	jwtCon := middleware.NewJWT()
	token, _ := jwtCon.CreateToken(middleware.CustomClaims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	})
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("username", username)

	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = "user"
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, err := client.GetRow(getRowRequest)
	if err != nil {
		fmt.Println("SearchUseamerIsExit  error:", err)
		return data, err
	}
	for _, v := range getResp.Columns {
		fmt.Println("v.ColumnName", v.ColumnName)
		if v.ColumnName == "password" {
			if v.Value == password {
				data = loginInfo{
					Uid:      getResp.Columns[2].Value.(int64),
					UserName: username,
					Status:   "success",
					Type:     getResp.Columns[1].Value.(string),
					Token:    token,
				}
				return data, nil
			}
		} else {
			return data, errors.New("密码错误")
		}
	}
	return data, errors.New("用户不存在")
}

//! 用户信息

type UserInfo struct {
	Uid      string `json:"uid"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Type     string `json:"type"`
	Desc     string `json:"desc"`
	Email    string `json:"email"`
}

//! 添加或修改用户信息
func HandleUserInfo(client *tablestore.TableStoreClient, data UserInfo) (interface{}, error) {
	info, err := SearchUserInfo(client, data.Uid, data.UserName)
	if err != nil {
		return nil, err
	}
	if info == nil && err == nil {
		res, err := SaveUserInfo(client, data)
		if err != nil {
			fmt.Println("SaveUserInfo:", err)
			return data, err
		} else {
			fmt.Println("SaveUserInfo finished", res)
		}
		return data, nil
	} else {
		res, err := updateUserInfo(client, data)
		if err != nil {
			fmt.Println("SaveUserInfo:", err)
			return data, err
		} else {
			fmt.Println("SaveUserInfo finished", res)
		}
		return data, nil
	}
}

func SaveUserInfo(client *tablestore.TableStoreClient, data UserInfo) (interface{}, error) {
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = allconst.Tables["userInfo"]
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("username", data.UserName)
	putPk.AddPrimaryKeyColumn("uid", data.Uid)
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("nickname", data.NickName)
	putRowChange.AddColumn("avatar", data.Avatar)
	putRowChange.AddColumn("type", data.Type)
	putRowChange.AddColumn("desc", data.Desc)
	putRowChange.AddColumn("email", data.Email)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	res, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("SaveUserInfo:", err)
		return data, err
	} else {
		fmt.Println("SaveUserInfo finished", res)
	}
	return data, nil
}

func updateUserInfo(client *tablestore.TableStoreClient, data UserInfo) (interface{}, error) {
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = allconst.Tables["userInfo"]
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("username", data.UserName)
	updatePk.AddPrimaryKeyColumn("uid", data.Uid)
	updateRowChange.PrimaryKey = updatePk
	updateRowChange.PutColumn("nickname", data.NickName)
	updateRowChange.PutColumn("avatar", data.Avatar)
	updateRowChange.PutColumn("type", data.Type)
	updateRowChange.PutColumn("desc", data.Desc)
	updateRowChange.PutColumn("email", data.Email)
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err := client.UpdateRow(updateRowRequest)

	if err != nil {
		fmt.Println("updateUserInfo:", err)
		return data, err
	} else {
		fmt.Println("update row finished")
	}
	return data, nil
}

//! 查询用户信息
func SearchUserInfo(client *tablestore.TableStoreClient, uid string, username string) (interface{}, error) {
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("username", username)
	putPk.AddPrimaryKeyColumn("uid", uid)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = allconst.Tables["userInfo"]
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, err := client.GetRow(getRowRequest)
	if err != nil {
		fmt.Println("SearchUseamerIsExit  error:", err)
		return nil, err
	}
	res := utils.HandleTableRes(getResp)
	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

//! 查询用户角色
func GetUserType(client *tablestore.TableStoreClient, username interface{}) (interface{}, error) {
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName("user")
	searchRequest.SetIndexName("user_index")
	query := &search.TermQuery{} //设置查询类型为TermQuery。
	query.FieldName = "username" //设置要匹配的字段。
	query.Term = username        //设置要匹配的值
	searchQuery := search.NewSearchQuery()
	searchQuery.SetQuery(query)
	searchQuery.SetGetTotalCount(true)
	searchRequest.SetSearchQuery(searchQuery)
	//设置为返回所有列。
	searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
		ReturnAll: true,
	})
	searchResponse, err := client.Search(searchRequest)
	userType := searchResponse.Rows[0].Columns[1].Value
	return userType, err
}
