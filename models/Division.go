package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Division struct {
	bun.BaseModel `bun:"divisions,select:divisions_view,alias:divisions_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Info          string        `json:"info"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`

	Address                 string      `json:"address"`
	Slug                    string      `json:"slug"`
	ShowCommonVisitingRules bool        `bun:"default:true" json:"showCommonVisitingRules"`
	Doctors                 Doctors     `bun:"rel:has-many" json:"doctors"`
	DoctorsForDelete        []uuid.UUID `bun:"-" json:"doctorsForDelete"`
	Vacancies               Vacancies   `bun:"rel:has-many" json:"vacancies"`
	Show                    bool        `json:"show"`

	DivisionPaidServices          DivisionPaidServices `bun:"rel:has-many" json:"divisionPaidServices"`
	DivisionPaidServicesForDelete []uuid.UUID          `bun:"-" json:"divisionPaidServicesForDelete"`

	Entrance                *Entrance        `bun:"rel:belongs-to" json:"entrance"`
	EntranceId              uuid.NullUUID    `bun:"type:uuid" json:"entranceId"`
	FloorId                 uuid.NullUUID    `bun:"type:uuid" json:"floorId"`
	Floor                   *Floor           `bun:"rel:belongs-to" json:"floor"`
	Timetable               *Timetable       `bun:"rel:belongs-to" json:"timetable"`
	TimetableId             uuid.UUID        `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
	Schedule                *Schedule        `bun:"rel:belongs-to" json:"schedule"`
	ScheduleId              uuid.UUID        `bun:"type:uuid" json:"scheduleId"`
	DivisionImages          DivisionImages   `bun:"rel:has-many" json:"divisionImages"`
	DivisionImagesForDelete []string         `bun:"-" json:"divisionImagesForDelete"`
	DivisionImagesNames     []string         `bun:"-" json:"divisionImagesNames"`
	DivisionComments        DivisionComments `bun:"rel:has-many" json:"divisionComments"`
	VisitingRules           VisitingRules    `bun:"rel:has-many" json:"visitingRules"`
	VisitingRulesForDelete  []uuid.UUID      `bun:"-" json:"visitingRulesForDelete"`

	HospitalizationContactInfoId uuid.UUID    `bun:"type:uuid" json:"hospitalizationContactInfoId"`
	HospitalizationContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"hospitalizationContactInfo"`

	HospitalizationDoctorID uuid.NullUUID `bun:"type:uuid" json:"hospitalizationDoctorId"`
	HospitalizationDoctor   *Doctor       `bun:"rel:belongs-to" json:"hospitalizationDoctor"`

	MedicalProfilesDivisions MedicalProfilesDivisions `bun:"rel:has-many" json:"medicalProfilesDivisions"`
	TreatDirection           *TreatDirection          `bun:"rel:belongs-to" json:"treatDirection"`
	TreatDirectionID         uuid.NullUUID            `bun:"type:uuid" json:"treatDirectionId"`
	Chief                    *Doctor                  `bun:"rel:belongs-to" json:"chief"`
	ChiefID                  uuid.NullUUID            `bun:"type:uuid" json:"chiefId"`
}

type Divisions []*Division

func (i Division) SetFilePath(fileID *string) *string {
	path := i.DivisionImages.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (i *Division) SetForeignKeys() {
	if i.HospitalizationContactInfo != nil {
		i.HospitalizationContactInfoId = i.HospitalizationContactInfo.ID
	}
	if i.ContactInfo != nil {
		i.ContactInfoID = i.ContactInfo.ID
	}
	if i.HospitalizationDoctor != nil {
		i.HospitalizationDoctorID = i.HospitalizationDoctor.ID
	}
	if i.Timetable != nil {
		i.TimetableId = i.Timetable.ID
	}
	if i.Schedule != nil {
		i.ScheduleId = i.Schedule.ID
	}
	if i.Chief != nil {
		i.ChiefID = i.Chief.ID
	}
}

func (items Divisions) GetSearchElements(searchGroup *SearchGroup) {
	//searchGroup.SearchElements = make(SearchElements, len(items))
	//for i := range items {
	//	searchGroup.SearchElements[i].Value = fmt.Sprintf("%s/%s", searchGroup.Prefix, items[i].Slug)
	//	searchGroup.SearchElements[i].Label = items[i].Name
	//}
}

func (item *Division) SetIdForChildren() {
	for i := range item.VisitingRules {
		item.VisitingRules[i].DivisionID = item.ID
	}
}
