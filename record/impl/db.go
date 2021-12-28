package record_impl

import "github.com/kldx/core/record"

type DB struct {}

func NewDB() record.IDB {
	return &DB{}
}

func (d *DB) Object(objectApiName string) record.IObject {
	return NewObject(objectApiName)
}

