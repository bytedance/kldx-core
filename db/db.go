package db

import "github.com/bytedance/kldx-core/structs"

type IDB interface {
	Object(objectApiName string) IObject
}

type IObject interface {
	// create
	Create(record interface{}) (*structs.RecordOnlyId, error)

	// update
	UpdateById(_id int64, record interface{}) error

	// delete
	DeleteById(_id int64) error

	// query
	Find(records interface{}) error
	FindOne(record interface{}) error

	// query related
	Where(condition interface{}) IQuery
	Offset(offset int64) IQuery
	Limit(limit int64) IQuery
	OrderBy(fieldApiNames ...string) IQuery
	OrderByDesc(fieldApiNames ...string) IQuery
	Select(fieldApiNames ...string) IQuery
}

type IQuery interface {
	// query
	Find(records interface{}) error
	FindOne(record interface{}) error

	// query related
	Where(condition interface{}) IQuery
	Offset(offset int64) IQuery
	Limit(limit int64) IQuery
	OrderBy(fieldApiNames ...string) IQuery
	OrderByDesc(fieldApiNames ...string) IQuery
	Select(fieldApiNames ...string) IQuery
}
