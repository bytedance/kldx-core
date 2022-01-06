package structs

type RecordOnlyId struct {
	Id int64 `json:"_id"`
}

type RecordsResult struct {
	Records []interface{} `json:"records"`
	Total   int64         `json:"total"`
}

type GetRecordsParam struct {
	Limit       int64         `json:"limit"`
	Offset      int64         `json:"offset"`
	Fields      []string      `json:"fields"`
	QuickSearch string        `json:"quickSearch"`
	Filter      []interface{} `json:"filter"`
	Sort        []*Sort       `json:"sort"`
}

type Sort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type FileUploadResult struct {
	FileId string `json:"fileId"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
}
