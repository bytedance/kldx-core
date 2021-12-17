package record

type CreateResult struct {
	Id int64 `json:"_id"`
}

type AsyncTaskResult struct {
	TaskId int64 `json:"taskID"`
}

type IObject interface {
	// 创建
	Create(record interface{}) (CreateResult, error)
	BatchCreate(records []interface{}) ([]int64, error)
	BatchCreateAsync(records []interface{}) (AsyncTaskResult, error)

	// 更新
	Update(record interface{}) error
	UpdateById(_id int64, record interface{}) error
	BatchUpdate(records []interface{}) error
	BatchUpdateAsync(records []interface{}) (AsyncTaskResult, error)

	// 删除
	Delete(record interface{}) error
	DeleteById(_id int64) error
	BatchDelete(records []interface{}) error
	BatchDeleteByIds(_ids []int64) error
	BatchDeleteAsync(records []interface{}) (AsyncTaskResult, error)
	BatchDeleteAsyncByIds(_ids []int64) (AsyncTaskResult, error)

	// 查询
	Find() ([]interface{}, error)
	FindOne() (interface{}, error)
	Count() (int64, error)

	// 查询相关
	Where() IQuery
	Offset(offset int64) IQuery
	Limit(offset int64) IQuery
	OrderBy(fieldApiNames... string) IQuery
	OrderByDesc(fieldApiNames... string) IQuery
	Select(fieldApiNames... string) IQuery
}

type IQuery interface {
	// 查询
	Find() ([]interface{}, error)
	FindOne() (interface{}, error)
	Count() (int64, error)

	// 查询相关
	Where() IQuery
	Offset(offset int64) IQuery
	Limit(offset int64) IQuery
	OrderBy(fieldApiNames... string) IQuery
	OrderByDesc(fieldApiNames... string) IQuery
	Select(fieldApiNames... string) IQuery
}
