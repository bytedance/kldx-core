package core

import (
	"code.byted.org/apaas/goapi_core/common/utils"
	"code.byted.org/apaas/goapi_core/db/field_type"
	"testing"
)

type User struct {
	ID        int64                    `json:"_id"`
	Name      *field_type.Multilingual `json:"_name"`
	IsDeleted bool                     `json:"_isDeleted"`
	Email     string                   `json:"_email,omitempty"`
}

func TestDB_User(t *testing.T) {
	var user User
	err := DB.Object("_user").Select("_name", "_email").FindOne(&user)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(user)
}

func TestDB_AllFieldsObject(t *testing.T) {
	var record AllFieldsObject
	err := DB.Object("allFieldsObject").Select("_id", "_createdBy", "_createdAt", "_updatedBy", "_updatedAt",
		"_name", "text", "number", "date", "datetime", "phone", "email", "option", "boolean", "multilingual", "richText",
		"attachment", "lookup", "autoid", "referenceField", "formula", "compositeType",
		//"avatar", TODO avatar 类型查询时会报系统错误
	).FindOne(&record)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(record)
}

func TestResouruces_File(t *testing.T) {
	res, err := Resources.File.UploadByPath("api_test.go", "./api_test.go")
	if err != nil {
		panic(err)
	}
	utils.PrintLog(res)

	data, err := Resources.File.Download(res.FileId)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(string(data))
}

type TestCompositeType struct {
	ID     int64  `json:"_id"`
	Number int64  `json:"number"`
	Text   string `json:"text"`
}

type AllFieldsObject struct {
	ID             int64                   `json:"_id"`
	Name           field_type.Multilingual `json:"_name"`
	Attachment     field_type.File         `json:"attachment"`
	Autoid         string                  `json:"autoid"`
	Boolean        bool                    `json:"boolean"`
	Date           string                  `json:"date"`
	Datetime       int64                   `json:"datetime"`
	Email          string                  `json:"email"`
	Formula        string                  `json:"formula"`
	Lookup         User                    `json:"lookup"`
	Multilingual   field_type.Multilingual `json:"multilingual"`
	Number         int64                   `json:"number"`
	Option         field_type.Option       `json:"option"`
	Phone          field_type.PhoneNumber  `json:"phone"`
	Text           string                  `json:"text"`
	ReferenceField field_type.Multilingual `json:"referenceField"`
	RichText       field_type.RichText     `json:"richText"`
	CompositeType  TestCompositeType       `json:"compositeType"`
	CreatedAt      int64                   `json:"_createdAt"`
	CreatedBy      User                    `json:"_createdBy"`
	UpdatedAt      int64                   `json:"_updatedAt"`
	UpdatedBy      User                    `json:"_updatedBy"`
	IsDeleted      bool                    `json:"_isDeleted"`
}
