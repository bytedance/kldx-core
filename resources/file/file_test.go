package file

import (
	"github.com/bytedance/kldx-core/common/utils"
	"os"
	"testing"
)

func TestFile_UploadByPath(t *testing.T) {
	result, err1 := NewFile().UploadByPath("a.txt", "./file_test.go")
	data, err2 := NewFile().Download(result.FileId)
	utils.PrintLog("TestFile_UploadByPath: ", err1, err2, result, string(data))
}

func TestFile_UploadByBuffer(t *testing.T) {
	buffer := []byte("这只是个测试")
	result, err1 := NewFile().UploadByBuffer("a.txt", buffer)
	data, err2 := NewFile().Download(result.FileId)
	utils.PrintLog("TestFile_UploadByBuffer: ", err1, err2, result, string(data))
}

func TestFile_UploadByReader(t *testing.T) {
	fileReader, err1 := os.Open("./file_test.go")
	result, err2 := NewFile().UploadByReader("a.txt", fileReader)
	data, err3 := NewFile().Download(result.FileId)
	utils.PrintLog("TestFile_UploadByReader: ", err1, err2, err3, result, string(data))
}

func TestFile_Download(t *testing.T) {
	data, err := NewFile().Download("eaeb81130a2a4581ad430170136f8689")
	utils.PrintLog("TestFile_Download: ", err, string(data))
}
