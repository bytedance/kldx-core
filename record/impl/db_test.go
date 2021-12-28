package record_impl

import (
	"github.com/kldx/core/common/utils"
	"github.com/kldx/core/record/condition"
	"testing"
)

func TestNewDB(t *testing.T) {
	var result interface{}
	err := NewDB().Object("employee").FindOne(&result)
	utils.PrintLog(err)
	utils.PrintLog(result)

	var results []interface{}
	err = NewDB().Object("employee").Select("name", "age").Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.And(
			cond.Eq("name", "小李"),
			cond.Gt("age", 0),
		)).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.Or(
			cond.Eq("name", "小李"),
			cond.Eq("name", "小花"),
		)).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	// 姓名是小刚或小花 同时 姓名是小刚或小明
	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.And(
			cond.Or(
				cond.Eq("name", "小刚"),
				cond.Eq("name", "小花"),
			),
			cond.Or(
				cond.Eq("name", "小刚"),
				cond.Eq("name", "小明"),
			))).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	// 查出 20 岁的小李 或 19岁的小花
	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.Or(cond.And(
			cond.Eq("name", "小李"),
			cond.Eq("age", 20),
		), cond.And(
			cond.Eq("name", "小花"),
			cond.Eq("age", 19),
		))).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	// 年龄是 19 或 20，姓名是 小李 或 小花
	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.And(cond.Or(
			cond.Eq("name", "小李"),
			cond.Eq("name", "小花"),
		), cond.Or(
			cond.Eq("age", 20),
			cond.Eq("age", 19),
		))).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	// 20 岁的小李和小花
	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.And(cond.Or(
			cond.Eq("name", "小李"),
			cond.Eq("name", "小花"),
		), cond.Eq("age", 20))).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)

	err = NewDB().Object("employee").Select("name", "age").Where(
		cond.And(cond.Or(
			cond.Eq("name", "小李"),
			cond.Eq("name", "小花"),
		), cond.Eq("age", 20))).Find(&results)
	utils.PrintLog(err)
	utils.PrintLog(results)
}
