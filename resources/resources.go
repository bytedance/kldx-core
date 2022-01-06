package resources

import "github.com/bytedance/kldx-core/resources/file"

type Resources struct {
	File file.IFile
}

func NewResources() *Resources {
	return &Resources{
		File: file.NewFile(),
	}
}
