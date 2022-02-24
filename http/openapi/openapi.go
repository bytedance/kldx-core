package openapi

import (
	"bytes"
	"encoding/json"
	cConstants "github.com/bytedance/kldx-common/constants"
	cExceptions "github.com/bytedance/kldx-common/exceptions"
	cHttp "github.com/bytedance/kldx-common/http"
	"github.com/bytedance/kldx-core/http"
	"github.com/bytedance/kldx-core/structs"
	"github.com/tidwall/gjson"
	"io"
	"mime/multipart"
	"path/filepath"
)

const (
	OpenapiSuccessCode_FileDownload = ""
	OpenapiSuccessCode_Success      = "0"

	OpenapiErrorCode_InternalError  = "k_ec_000001"
	OpenapiErrorCode_NoTenantID     = "k_ec_000002"
	OpenapiErrorCode_NoUserID       = "k_ec_000003"
	OpenapiErrorCode_UnknownError   = "k_ec_000004"
	OpenapiErrorCode_OpUnknownError = "k_op_ec_00001"
	OpenapiErrorCode_SystemBusy     = "k_op_ec_20001"
	OpenapiErrorCode_SystemError    = "k_op_ec_20002"
	OpenapiErrorCode_RateLimitError = "k_op_ec_20003"
	OpenapiErrorCode_TokenExpire    = "k_ident_013000"
	OpenapiErrorCode_IllegalToken   = "k_ident_013001"
	OpenapiErrorCode_MissingToken   = "k_op_ec_10205"
)

func errorWrapper(body []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, cExceptions.ErrorWrap(err)
	}

	code := gjson.GetBytes(body, "code").String()
	msg := gjson.GetBytes(body, "msg").String()
	switch code {
	case OpenapiSuccessCode_FileDownload:
		return body, nil
	case OpenapiSuccessCode_Success:
		data := gjson.GetBytes(body, "data")
		if data.Type == gjson.String {
			return []byte(data.Str), nil
		}
		return []byte(data.Raw), nil
	case OpenapiErrorCode_InternalError, OpenapiErrorCode_NoTenantID, OpenapiErrorCode_NoUserID, OpenapiErrorCode_UnknownError,
		OpenapiErrorCode_OpUnknownError, OpenapiErrorCode_SystemBusy, OpenapiErrorCode_SystemError, OpenapiErrorCode_RateLimitError,
		OpenapiErrorCode_TokenExpire, OpenapiErrorCode_IllegalToken, OpenapiErrorCode_MissingToken:
		return nil, cExceptions.InternalError("request openapi failed, code: %s, msg: %s", code, msg)
	default:
		return nil, cExceptions.InvalidParamError("request openapi failed, code: %s, msg: %s", code, msg)
	}
}

func CreateRecord(objectApiName string, record interface{}) (*structs.RecordOnlyId, error) {
	data, err := errorWrapper(http.GetOpenapiClient().PostJson(http.GetOpenapiPath_CreateRecord(objectApiName), nil, record, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	result := structs.RecordOnlyId{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("CreateRecord failed, err: %v", err)
	}

	return &result, nil
}

func UpdateRecord(objectApiName string, recordId int64, record interface{}) (*structs.RecordOnlyId, error) {
	data, err := errorWrapper(http.GetOpenapiClient().PatchJson(http.GetOpenapiPath_UpdateRecord(objectApiName, recordId), nil, record, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	result := structs.RecordOnlyId{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("UpdateRecord failed, err: %v", err)
	}

	return &result, nil
}

func DeleteRecordById(objectApiName string, recordId int64) error {
	_, err := errorWrapper(http.GetOpenapiClient().DeleteJson(http.GetOpenapiPath_DeleteRecordById(objectApiName, recordId), nil, nil, cHttp.AppTokenMiddleware))
	return err
}

func GetRecordById(objectApiName string, recordId int64) (interface{}, error) {
	data, err := errorWrapper(http.GetOpenapiClient().Get(http.GetOpenapiPath_GetRecordById(objectApiName, recordId), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("GetRecordById failed, err: %v", err)
	}

	return &result, nil
}

func GetRecords(objectApiName string, param *structs.GetRecordsParam) ([]interface{}, error) {
	if param == nil {
		param = &structs.GetRecordsParam{}
	}

	if param.Limit == 0 {
		param.Limit = 200
	}

	data, err := errorWrapper(http.GetOpenapiClient().PostJson(http.GetOpenapiPath_GetRecords(objectApiName), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	result := structs.RecordsResult{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("GetRecords failed, err: %v", err)
	}

	return result.Records, nil
}

func UploadFile(fileName string, fileReader io.Reader) (*structs.FileUploadResult, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	form, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	_, err = io.Copy(form, fileReader)
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	headers := map[string][]string{
		cConstants.HttpHeaderKey_ContentType: {writer.FormDataContentType()},
	}

	data, err := errorWrapper(http.GetOpenapiClient().PostFormData(http.OpenapiPath_UploadFile, headers, payload, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var result structs.FileUploadResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	return &result, nil
}

func DownloadFile(fileId string) ([]byte, error) {
	data, err := errorWrapper(http.GetOpenapiClient().Get(http.GetOpenapiPath_DownloadFile(fileId), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	return data, nil
}