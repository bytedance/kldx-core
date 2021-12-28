package openapi

import (
	"github.com/kldx/core/structs"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateRecord(t *testing.T) {
	result, err := CreateRecord("employee", map[string]interface{}{
		"name": "小花",
		"age":  20,
	})
	fmt.Println(err)
	fmt.Println(result)
}

func TestUpdateRecordById(t *testing.T) {
	result, err := UpdateRecord("employee", 1719656278569063, map[string]interface{}{
		"name": "小花",
		"age":  19,
	})
	fmt.Println(err)
	fmt.Println(result)
}

func TestDeleteRecordById(t *testing.T) {
	result, err := CreateRecord("employee", map[string]interface{}{
		"name": "小明",
		"age":  18,
	})
	fmt.Println(err)
	fmt.Println(result)

	err = DeleteRecordById("employee", result.Id)
	fmt.Println(err)
}

func TestGetRecordById(t *testing.T) {
	result, err := GetRecordById("employee", 1719656278569063)
	res, _ := json.Marshal(result)
	fmt.Println(err)
	fmt.Println(string(res))
}

func TestGetRecords(t *testing.T) {
	result, err := GetRecords("employee", &structs.GetRecordsParam{
		Limit: 300,
	})
	fmt.Println(err)
	fmt.Println(result)
}
