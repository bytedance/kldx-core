package resources

import "code.byted.org/apaas/goapi_core/resources/file"

type Resources struct {
	File file.IFile
}

func NewResources() *Resources {
	return &Resources{
		File: file.NewFile(),
	}
}
