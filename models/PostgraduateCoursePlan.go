package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type PostgraduateCoursePlan struct {
	bun.BaseModel `bun:"postgraduate_course_plans,alias:postgraduate_course_plans"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Year          time.Time     `json:"year"`
	Plan          *FileInfo     `bun:"rel:belongs-to" json:"plan"`
	PlanID        uuid.NullUUID `bun:"type:uuid" json:"planId"`

	PostgraduateCourse   *PostgraduateCourse `bun:"rel:belongs-to" json:"postgraduateCourse"`
	PostgraduateCourseID uuid.NullUUID       `bun:"type:uuid" json:"postgraduateCourseId"`
}

type PostgraduateCoursePlans []*PostgraduateCoursePlan

func (items PostgraduateCoursePlans) GetPlans() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].Plan)
	}
	return itemsForGet
}

func (item *PostgraduateCoursePlan) SetForeignKeys() {
	item.PlanID = item.Plan.ID
}

func (items PostgraduateCoursePlans) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (items PostgraduateCoursePlans) SetFilePath(fileID string) *string {
	for i := range items {
		filePath := items[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *PostgraduateCoursePlan) SetFilePath(fileID string) *string {
	if item.Plan.ID.UUID.String() == fileID {
		item.Plan.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Plan.FileSystemPath
	}
	return nil
}
