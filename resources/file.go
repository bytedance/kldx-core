package resources

import (
	"bytes"
	cExceptions "github.com/kldx/common/exceptions"
	"github.com/kldx/core/http/openapi"
	"github.com/kldx/core/structs"
	"io"
	"os"
)

type IFile interface {
	UploadByPath(fileName, path string) (*structs.FileUploadResult, error)
	UploadByReader(fileName string, reader io.Reader) (*structs.FileUploadResult, error)
	UploadByBuffer(fileName string, buffer []byte) (*structs.FileUploadResult, error)
	Download(fileId string) ([]byte, error)
}

type File struct {}

func NewFile() *File {
	return &File{}
}

func (f *File) UploadByPath(fileName, path string) (*structs.FileUploadResult, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, cExceptions.InvalidParamError("UploadByPath failed, err: %v", err)
	}
	return openapi.UploadFile(fileName, fileReader)
}

func (f *File) UploadByReader(fileName string, reader io.Reader) (*structs.FileUploadResult, error) {
	return openapi.UploadFile(fileName, reader)
}

func (f *File) UploadByBuffer(fileName string, buffer []byte) (*structs.FileUploadResult, error) {
	return openapi.UploadFile(fileName, bytes.NewReader(buffer))
}

func (f *File) Download(fileId string) ([]byte, error) {
	return openapi.DownloadFile(fileId)
}

