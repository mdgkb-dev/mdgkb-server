package models

import "github.com/uptrace/bun"

type FormPattern struct {
	bun.BaseModel `bun:"form_patterns,alias:form_patterns"`
	Title         string `json:"title"`
	Form
}

type FormPatterns []*FormPattern

func (item *FormPattern) SetIdForChildren() {
	if len(item.Fields) == 0 {
		return
	}
	for i := range item.Fields {
		item.Fields[i].FormPatternID = item.ID
	}
}

func (item *FormPattern) SetFilePath(fileID *string) *string {
	for i := range item.Fields {
		filePath := item.Fields[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}
