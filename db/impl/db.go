package record_impl

import (
	"github.com/bytedance/kldx-core/db"
)

type DB struct{}

func NewDB() db.IDB {
	return &DB{}
}

func (d *DB) Object(objectApiName string) db.IObject {
	return NewObject(objectApiName)
}
