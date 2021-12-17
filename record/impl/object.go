package record_impl

import (
	"code.byted.org/zhouwexin/kldx_core/record"
)

func newObject(objectApiName string) *Object {
	return &Object{
		objectApiName: objectApiName,
	}
}

type Object struct {
	objectApiName string
}

func (o *Object) Create(record interface{}) (record.CreateResult, error) {
	panic("implement me")
}

func (o *Object) BatchCreate(records []interface{}) ([]int64, error) {
	panic("implement me")
}

func (o *Object) BatchCreateAsync(records []interface{}) (record.AsyncTaskResult, error) {
	panic("implement me")
}

func (o *Object) Update(record interface{}) error {
	panic("implement me")
}

func (o *Object) UpdateById(_id int64, record interface{}) error {
	panic("implement me")
}

func (o *Object) BatchUpdate(records []interface{}) error {
	panic("implement me")
}

func (o *Object) BatchUpdateAsync(records []interface{}) (record.AsyncTaskResult, error) {
	panic("implement me")
}

func (o *Object) Delete(record interface{}) error {
	panic("implement me")
}

func (o *Object) DeleteById(_id int64) error {
	panic("implement me")
}

func (o *Object) BatchDelete(records []interface{}) error {
	panic("implement me")
}

func (o *Object) BatchDeleteByIds(_ids []int64) error {
	panic("implement me")
}

func (o *Object) BatchDeleteAsync(records []interface{}) (record.AsyncTaskResult, error) {
	panic("implement me")
}

func (o *Object) BatchDeleteAsyncByIds(_ids []int64) (record.AsyncTaskResult, error) {
	panic("implement me")
}

func (o *Object) Find() ([]interface{}, error) {
	panic("implement me")
}

func (o *Object) FindOne() (interface{}, error) {
	panic("implement me")
}

func (o *Object) Count() (int64, error) {
	panic("implement me")
}

func (o *Object) Where() record.IQuery {
	panic("implement me")
}

func (o *Object) Offset(offset int64) record.IQuery {
	panic("implement me")
}

func (o *Object) Limit(offset int64) record.IQuery {
	panic("implement me")
}

func (o *Object) OrderBy(fieldApiNames ...string) record.IQuery {
	panic("implement me")
}

func (o *Object) OrderByDesc(fieldApiNames ...string) record.IQuery {
	panic("implement me")
}

func (o *Object) Select(fieldApiNames ...string) record.IQuery {
	panic("implement me")
}
