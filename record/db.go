package record

type IDB interface {
	Object(objectApiName string) IObject
}

