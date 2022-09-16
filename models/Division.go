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
	EntranceID              uuid.NullUUID    `bun:"type:uuid" json:"entranceId"`
	FloorID                 uuid.NullUUID    `bun:"type:uuid" json:"floorId"`
	Floor                   *Floor           `bun:"rel:belongs-to" json:"floor"`
	Timetable               *Timetable       `bun:"rel:belongs-to" json:"timetable"`
	TimetableID             uuid.NullUUID    `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
	Schedule                *Schedule        `bun:"rel:belongs-to" json:"schedule"`
	ScheduleID              uuid.UUID        `bun:"type:uuid" json:"scheduleId"`
	DivisionImages          DivisionImages   `bun:"rel:has-many" json:"divisionImages"`
	DivisionImagesForDelete []string         `bun:"-" json:"divisionImagesForDelete"`
	DivisionImagesNames     []string         `bun:"-" json:"divisionImagesNames"`
	DivisionComments        DivisionComments `bun:"rel:has-many" json:"divisionComments"`
	VisitingRules           VisitingRules    `bun:"rel:has-many" json:"visitingRules"`
	VisitingRulesForDelete  []uuid.UUID      `bun:"-" json:"visitingRulesForDelete"`

	HospitalizationContactInfoID uuid.UUID    `bun:"type:uuid" json:"hospitalizationContactInfoId"`
	HospitalizationContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"hospitalizationContactInfo"`

	HospitalizationDoctorID uuid.NullUUID `bun:"type:uuid" json:"hospitalizationDoctorId"`
	HospitalizationDoctor   *Doctor       `bun:"rel:belongs-to" json:"hospitalizationDoctor"`

	MedicalProfilesDivisions MedicalProfilesDivisions `bun:"rel:has-many" json:"medicalProfilesDivisions"`
	TreatDirection           *TreatDirection          `bun:"rel:belongs-to" json:"treatDirection"`
	TreatDirectionID         uuid.NullUUID            `bun:"type:uuid" json:"treatDirectionId"`
	Chief                    *Doctor                  `bun:"rel:belongs-to" json:"chief"`
	ChiefID                  uuid.NullUUID            `bun:"type:uuid" json:"chiefId"`

	NewsDivisions          NewsDivisions `bun:"rel:has-many" json:"newsDivisions"`
	NewsDivisionsForDelete []uuid.UUID   `bun:"-" json:"newsDivisionsForDelete"`
}

type Divisions []*Division

func (item *Division) SetFilePath(fileID *string) *string {
	path := item.DivisionImages.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *Division) SetForeignKeys() {
	if item.HospitalizationContactInfo != nil {
		item.HospitalizationContactInfoID = item.HospitalizationContactInfo.ID
	}
	if item.ContactInfo != nil {
		item.ContactInfoID = item.ContactInfo.ID
	}
	if item.HospitalizationDoctor != nil {
		item.HospitalizationDoctorID = item.HospitalizationDoctor.ID
	}
	if item.Timetable != nil {
		item.TimetableID = item.Timetable.ID
	}
	if item.Schedule != nil {
		item.ScheduleID = item.Schedule.ID
	}
	if item.Chief != nil {
		item.ChiefID = item.Chief.ID
	}
	if item.TreatDirection != nil {
		item.TreatDirectionID = item.TreatDirection.ID
	}
}

func (items Divisions) GetSearchElements(searchGroup *SearchGroup) {
	//searchGroup.SearchElements = make(SearchElements, len(items))
	//for i := range items {
	//	searchGroup.SearchElements[i].Value = fmt.Sprintf("%s/%s", searchGroup.Prefix, items[i].Slug)
	//	searchGroup.SearchElements[i].Label = items[i].Name
	//}
}

func (item *Division) SetIDForChildren() {
	for i := range item.VisitingRules {
		item.VisitingRules[i].DivisionID = item.ID
	}
	for i := range item.NewsDivisions {
		item.NewsDivisions[i].DivisionID = item.ID
	}
	for i := range item.MedicalProfilesDivisions {
		item.MedicalProfilesDivisions[i].DivisionID = item.ID
	}

}
