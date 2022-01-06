package record_impl

import (
	"code.byted.org/apaas/goapi_core/db"
)

type DB struct{}

func NewDB() db.IDB {
	return &DB{}
}

func (d *DB) Object(objectApiName string) db.IObject {
	return NewObject(objectApiName)
}
