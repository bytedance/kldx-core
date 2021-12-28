package record_impl

import (
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_core/common/constants"
	"code.byted.org/apaas/goapi_core/http/openapi"
	"code.byted.org/apaas/goapi_core/record"
	op "code.byted.org/apaas/goapi_core/record/operator"
	"code.byted.org/apaas/goapi_core/structs"
	"encoding/json"
	"reflect"
)

type Object struct {
	objectApiName string
	err           error
}

func NewObject(objectApiName string) *Object {
	o := &Object{
		objectApiName: objectApiName,
	}

	if objectApiName == "" {
		o.err = cExceptions.InvalidParamError("[Object] objectApiName is empty")
	}
	return o
}

func (o *Object) Create(record interface{}) (*structs.RecordOnlyId, error) {
	if o.err != nil {
		return nil, o.err
	}
	return openapi.CreateRecord(o.objectApiName, record)
}

func (o *Object) UpdateById(_id int64, record interface{}) error {
	if o.err != nil {
		return o.err
	}
	_, err := openapi.UpdateRecord(o.objectApiName, _id, record)
	return err
}

func (o *Object) DeleteById(_id int64) error {
	if o.err != nil {
		return o.err
	}
	return openapi.DeleteRecordById(o.objectApiName, _id)
}

func (o *Object) Find(records interface{}) error {
	if o.err != nil {
		return o.err
	}
	return newQuery(o.objectApiName, o.err).Find(records)
}

func (o *Object) FindOne(record interface{}) error {
	if o.err != nil {
		return o.err
	}
	return newQuery(o.objectApiName, o.err).FindOne(record)
}

func (o *Object) Where(condition interface{}) record.IQuery {
	return newQuery(o.objectApiName, o.err).Where(condition)
}

func (o *Object) Offset(offset int64) record.IQuery {
	return newQuery(o.objectApiName, o.err).Offset(offset)
}

func (o *Object) Limit(limit int64) record.IQuery {
	return newQuery(o.objectApiName, o.err).Limit(limit)
}

func (o *Object) OrderBy(fieldApiNames ...string) record.IQuery {
	return newQuery(o.objectApiName, o.err).OrderBy(fieldApiNames...)
}

func (o *Object) OrderByDesc(fieldApiNames ...string) record.IQuery {
	return newQuery(o.objectApiName, o.err).OrderByDesc(fieldApiNames...)
}

func (o *Object) Select(fieldApiNames ...string) record.IQuery {
	return newQuery(o.objectApiName, o.err).Select(fieldApiNames...)
}

type Query struct {
	objectApiName string
	limit         int64
	offset        int64
	fields        []string
	order         []*structs.Sort
	conditions    []interface{}
	err           error
}

func newQuery(objectApiName string, err error) *Query {
	q := &Query{
		objectApiName: objectApiName,
		limit:         200,
		offset:        0,
		fields:        []string{},
		order:         []*structs.Sort{},
	}

	if err != nil {
		q.err = err
		return q
	}

	if objectApiName == "" {
		q.err = cExceptions.InvalidParamError("objectApiName is empty")
		return q
	}

	return q
}

func (q *Query) Find(records interface{}) error {
	if q.err != nil {
		return q.err
	}

	param := &structs.GetRecordsParam{
		Limit:  q.limit,
		Offset: q.offset,
		Fields: q.fields,
		Filter: q.conditions,
		Sort:   q.order,
	}

	originRecords, err := openapi.GetRecords(q.objectApiName, param)
	if err != nil {
		return err
	}
	recordsStr, err := json.Marshal(originRecords)
	if err != nil {
		return cExceptions.InternalError("Query.Find failed, err: %v", err)
	}

	return json.Unmarshal(recordsStr, &records)
}

func (q *Query) FindOne(record interface{}) error {
	if q.err != nil {
		return q.err
	}

	param := &structs.GetRecordsParam{
		Limit:  1,
		Offset: q.offset,
		Fields: q.fields,
		Filter: q.conditions,
		Sort:   q.order,
	}

	originRecords, err := openapi.GetRecords(q.objectApiName, param)
	if err != nil {
		return err
	}

	if len(originRecords) == 0 {
		return nil
	}

	recordStr, err := json.Marshal(originRecords[0])
	if err != nil {
		return cExceptions.InternalError("Query.Find failed, err: %v", err)
	}

	return json.Unmarshal(recordStr, &record)
}

func (q *Query) Where(condition interface{}) record.IQuery {
	if q.err != nil {
		return q
	}

	if condition == nil {
		return q
	}

	typ := reflect.TypeOf(condition)
	val := reflect.ValueOf(condition)
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	switch typ.Kind() {
	case reflect.Slice:
		q.conditions = append(q.conditions, map[string]interface{}{op.And: condition})
	case reflect.Struct, reflect.Map:
		q.conditions = append(q.conditions, condition)
	default:
		q.err = cExceptions.InvalidParamError("Query.Where received invalid type, should be slice, struct or map, but received %s ", typ)
	}
	return q
}

func (q *Query) Offset(offset int64) record.IQuery {
	if q.err != nil {
		return q
	}

	if offset < 0 {
		q.err = cExceptions.InvalidParamError("Query.Offset received invalid value, should >= 0")
	}

	q.offset = offset
	return q
}

func (q *Query) Limit(limit int64) record.IQuery {
	if q.err != nil {
		return q
	}

	if limit < 1 || limit > 200 {
		q.err = cExceptions.InvalidParamError("Query.Limit received invalid value (%d), should be 1~200", limit)
	}

	q.limit = limit
	return q
}

func (q *Query) OrderBy(fieldApiNames ...string) record.IQuery {
	if q.err != nil {
		return q
	}

	for _, fieldApiName := range fieldApiNames {
		q.order = append(q.order, &structs.Sort{
			Field:     fieldApiName,
			Direction: constants.OrderAsc,
		})
	}
	return q
}

func (q *Query) OrderByDesc(fieldApiNames ...string) record.IQuery {
	if q.err != nil {
		return q
	}

	for _, fieldApiName := range fieldApiNames {
		q.order = append(q.order, &structs.Sort{
			Field:     fieldApiName,
			Direction: constants.OrderDesc,
		})
	}
	return q
}

func (q *Query) Select(fieldApiNames ...string) record.IQuery {
	if q.err != nil {
		return q
	}

	q.fields = append(q.fields, fieldApiNames...)
	return q
}
