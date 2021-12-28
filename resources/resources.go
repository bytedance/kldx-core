package resources

type Resources struct {
	File IFile
}

func NewResources() *Resources {
	return &Resources{
		File: NewFile(),
	}
}
