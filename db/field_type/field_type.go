package field_type

type Multilingual struct {
	En string `json:"en_US,omitempty"`
	Zh string `json:"zh_CN,omitempty"`
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	DialingCode string `json:"dialingCode"`
	Number      string `json:"number"`
}

type Option struct {
	ApiName string       `json:"apiName"`
	Label   Multilingual `json:"label"`
}

type File struct {
	FileId string `json:"fileId"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Type   string `json:"type"`
}

type RichTextConfig struct {
	ResourceId   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
}

type RichText struct {
	Config []*RichTextConfig `json:"config"`
	Raw    string            `json:"raw"`
}
