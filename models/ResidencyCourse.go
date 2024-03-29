package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type ResidencyCourse struct {
	bun.BaseModel                            `bun:"residency_courses,select:residency_courses_view,alias:residency_courses_view"`
	ID                                       uuid.NullUUID                   `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Description                              string                          `json:"description"`
	EducationForm                            string                          `json:"educationForm"`
	Slug                                     string                          `bun:",scanonly" json:"slug"`
	FreePlaces                               int                             `json:"freePlaces"`
	FreeGovernmentPlaces                     int                             `json:"freeGovernmentPlaces"`
	Cost                                     int                             `json:"cost"`
	PaidPlaces                               int                             `json:"paidPlaces"`
	ResidencyCoursesSpecializations          ResidencyCoursesSpecializations `bun:"rel:has-many" json:"residencyCoursesSpecializations"`
	ResidencyCoursesSpecializationsForDelete []uuid.UUID                     `bun:"-" json:"residencyCoursesSpecializationsForDelete"`

	ResidencyApplications ResidencyApplications `bun:"rel:has-many" json:"residencyApplications"`

	StartYear   *EducationYear `bun:"rel:belongs-to" json:"startYear"`
	StartYearID uuid.NullUUID  `bun:"type:uuid" json:"startYearId"`

	EndYear   *EducationYear `bun:"rel:belongs-to" json:"endYear"`
	EndYearID uuid.NullUUID  `bun:"type:uuid" json:"endYearId"`

	Annotation   *FileInfo     `bun:"rel:belongs-to" json:"annotation"`
	AnnotationID uuid.NullUUID `bun:"type:uuid" json:"annotationId"`

	Program   *FileInfo     `bun:"rel:belongs-to" json:"program"`
	ProgramID uuid.NullUUID `bun:"type:uuid" json:"programId"`

	Plan   *FileInfo     `bun:"rel:belongs-to" json:"plan"`
	PlanID uuid.NullUUID `bun:"type:uuid" json:"planId"`

	Schedule   *FileInfo     `bun:"rel:belongs-to" json:"schedule"`
	ScheduleID uuid.NullUUID `bun:"type:uuid" json:"scheduleId"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`

	ResidencyCoursePracticePlaceGroups          ResidencyCoursePracticePlaceGroups `bun:"rel:has-many"  json:"residencyCoursePracticePlaceGroups"`
	ResidencyCoursePracticePlaceGroupsForDelete []uuid.UUID                        `bun:"-"  json:"residencyCoursePracticePlaceGroupsForDelete"`

	Name          string        `bun:"-" json:"name"`
	MainTeacher   *Employee     `bun:"rel:belongs-to" json:"mainTeacher"`
	MainTeacherID uuid.NullUUID `bun:"type:uuid" json:"mainTeacherId,omitempty"`
}

type ResidencyCourses []*ResidencyCourse

func (item *ResidencyCourse) SetForeignKeys() {
	item.AnnotationID = item.Annotation.ID
	item.ProgramID = item.Program.ID
	item.PlanID = item.Plan.ID
	item.FormPatternID = item.FormPattern.ID
	item.ScheduleID = item.Schedule.ID
}

func (item *ResidencyCourse) SetIDForChildren() {
	for i := range item.ResidencyCoursesSpecializations {
		item.ResidencyCoursesSpecializations[i].ResidencyCourseID = item.ID
	}
	for i := range item.ResidencyCoursePracticePlaceGroups {
		item.ResidencyCoursePracticePlaceGroups[i].ResidencyCourseID = item.ID
	}
}

func (item *ResidencyCourse) SetFilePath(fileID string) *string {
	if item.Annotation.ID.UUID.String() == fileID {
		item.Annotation.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Annotation.FileSystemPath
	}
	if item.Plan.ID.UUID.String() == fileID {
		item.Plan.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Plan.FileSystemPath
	}
	if item.Program.ID.UUID.String() == fileID {
		item.Program.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Program.FileSystemPath
	}
	if item.Schedule.ID.UUID.String() == fileID {
		item.Schedule.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Schedule.FileSystemPath
	}
	return nil
}
