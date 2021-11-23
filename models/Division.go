package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Division struct {
	bun.BaseModel `bun:"divisions,alias:divisions"`
	ID                      uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                    string           `json:"name"`
	Info                    string           `json:"info"`
	Phone                   string           `json:"phone"`
	Email                   string           `json:"email"`
	Address                 string           `json:"address"`
	Slug                    string           `json:"slug"`
	Doctors                 Doctors          `bun:"rel:has-many" json:"doctors"`
	Vacancies               Vacancies        `bun:"rel:has-many" json:"vacancies"`
	Entrance                *Entrance        `bun:"rel:belongs-to" json:"entrance"`
	EntranceId              uuid.UUID        `bun:"type:uuid" json:"entranceId"`
	FloorId                 uuid.UUID        `bun:"type:uuid" json:"floorId"`
	Timetable               *Timetable       `bun:"rel:belongs-to" json:"timetable"`
	TimetableId             uuid.UUID        `bun:"type:uuid" json:"timetableId"`
	Schedule                *Schedule        `bun:"rel:belongs-to" json:"schedule"`
	ScheduleId              uuid.UUID        `bun:"type:uuid" json:"scheduleId"`
	DivisionImages          DivisionImages   `bun:"rel:has-many" json:"divisionImages"`
	DivisionImagesForDelete []string         `bun:"-" json:"divisionImagesForDelete"`
	DivisionImagesNames     []string         `bun:"-" json:"divisionImagesNames"`
	DivisionComments        DivisionComments `bun:"rel:has-many" json:"divisionComments"`
}

type Divisions []*Division

func (i Division) SetFilePath(fileID *string) *string {
	path := i.DivisionImages.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items Divisions) GetSearchElements(searchGroup *SearchGroup) {
	//searchGroup.SearchElements = make(SearchElements, len(items))
	//for i := range items {
	//	searchGroup.SearchElements[i].Value = fmt.Sprintf("%s/%s", searchGroup.Prefix, items[i].Slug)
	//	searchGroup.SearchElements[i].Label = items[i].Name
	//}
}
