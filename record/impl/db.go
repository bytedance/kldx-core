package record_impl

import "code.byted.org/apaas/goapi_core/record"

type DB struct {}

func NewDB() record.IDB {
	return &DB{}
}

func (d *DB) Object(objectApiName string) record.IObject {
	return NewObject(objectApiName)
}

