package mock

import (
	"encoding/json"
	"io/ioutil"
	"os"

	querypb "github.com/xelabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

func ConnectionID() *sqltypes.Result {
	return &sqltypes.Result{
		Fields: []*querypb.Field{
			{
				Name: "CONNECTION_ID()",
				Type: querypb.Type_INT32,
			},
		},
		Rows: [][]sqltypes.Value{
			{
				sqltypes.MakeTrusted(querypb.Type_INT32, []byte("10")),
			},
		},
	}
}

func TableRes() *sqltypes.Result {

	return &sqltypes.Result{
		Fields: []*querypb.Field{
			{
				Name: "id",
				Type: querypb.Type_INT32,
			},
			{
				Name: "name",
				Type: querypb.Type_VARCHAR,
			},
		},
		Rows: [][]sqltypes.Value{
			{
				sqltypes.MakeTrusted(querypb.Type_INT32, []byte("10")),
				sqltypes.MakeTrusted(querypb.Type_VARCHAR, []byte("nice name")),
			},
		},
	}

}

func read(filename string) []byte {
	// 打开 JSON 文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件内容
	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

type ShowV struct {
	Table string `json:"table"`
	Rows  []struct {
		VariableName string `json:"Variable_name"`
		Value        string `json:"Value"`
	} `json:"rows"`
}

func ShowVal() *sqltypes.Result {

	var showV ShowV

	// 解析 JSON 数据到 User 实例
	err := json.Unmarshal(read("showval.json"), &showV)
	if err != nil {
		panic(err)
	}

	var rows [][]sqltypes.Value

	// 确保rows有足够的行，并且每一行都初始化为一个空切片
	for i := range showV.Rows {
		println(i)
		rows = append(rows, []sqltypes.Value{}) // 初始化每一行为空切片
	}

	for k, v := range showV.Rows {
		rows[k] = append(rows[k], sqltypes.MakeTrusted(querypb.Type_VARCHAR, []byte(v.VariableName)))
		rows[k] = append(rows[k], sqltypes.MakeTrusted(querypb.Type_VARCHAR, []byte(v.Value)))
	}

	return &sqltypes.Result{
		Fields: []*querypb.Field{
			{
				Name: "Variable_name",
				Type: querypb.Type_VARCHAR,
			},
			{
				Name: "Value",
				Type: querypb.Type_VARCHAR,
			},
		},
		Rows: rows,
	}
}