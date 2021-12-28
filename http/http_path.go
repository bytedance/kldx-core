package http

import (
	cConstants "code.byted.org/apaas/goapi_common/constants"
	cUtils "code.byted.org/apaas/goapi_common/utils"
	"strconv"
	"strings"
)

const (
	OpenapiPath_CreateRecord     = "/api/data/v1/namespaces/:namespace/objects/:objectApiName"
	OpenapiPath_UpdateRecordById = "/api/data/v1/namespaces/:namespace/objects/:objectApiName/:recordId"
	OpenapiPath_GetRecordById    = "/api/data/v1/namespaces/:namespace/objects/:objectApiName/:recordId"
	OpenapiPath_DeleteRecordById = "/api/data/v1/namespaces/:namespace/objects/:objectApiName/:recordId"
	OpenapiPath_GetRecords       = "/api/data/v1/namespaces/:namespace/objects/:objectApiName/records"
	OpenapiPath_UploadFile       = "/api/attachment/v1/files"
	OpenapiPath_DownloadFile     = "/api/attachment/v1/files/:fileId"
)

func GetOpenapiPath_CreateRecord(objectApiName string) string {
	path := strings.ReplaceAll(OpenapiPath_CreateRecord, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	return strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
}

func GetOpenapiPath_UpdateRecord(objectApiName string, recordId int64) string {
	path := strings.ReplaceAll(OpenapiPath_UpdateRecordById, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	path = strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
	return strings.ReplaceAll(path, cConstants.ReplaceRecordId, strconv.FormatInt(recordId, 10))
}

func GetOpenapiPath_GetRecordById(objectApiName string, recordId int64) string {
	path := strings.ReplaceAll(OpenapiPath_GetRecordById, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	path = strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
	return strings.ReplaceAll(path, cConstants.ReplaceRecordId, strconv.FormatInt(recordId, 10))
}

func GetOpenapiPath_DeleteRecordById(objectApiName string, recordId int64) string {
	path := strings.ReplaceAll(OpenapiPath_DeleteRecordById, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	path = strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
	return strings.ReplaceAll(path, cConstants.ReplaceRecordId, strconv.FormatInt(recordId, 10))
}

func GetOpenapiPath_GetRecords(objectApiName string) string {
	path := strings.ReplaceAll(OpenapiPath_GetRecords, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	return strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
}

func GetOpenapiPath_DownloadFile(fileId string) string {
	return strings.ReplaceAll(OpenapiPath_DownloadFile, cConstants.ReplaceFileId, fileId)
}
