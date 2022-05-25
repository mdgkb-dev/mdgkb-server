package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
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
	ResidencyCoursesTeachers                 ResidencyCoursesTeachers        `bun:"rel:has-many" json:"residencyCoursesTeachers"`
	ResidencyCoursesTeachersForDelete        []uuid.UUID                     `bun:"-" json:"residencyCoursesForDelete"`

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
}

type ResidencyCourses []*ResidencyCourse

func (item *ResidencyCourse) SetForeignKeys() {
	item.AnnotationID = item.Annotation.ID
	item.ProgramID = item.Program.ID
	item.PlanID = item.Plan.ID
	item.FormPatternID = item.FormPattern.ID
	item.ScheduleID = item.Schedule.ID
}

func (item *ResidencyCourse) SetIdForChildren() {
	for i := range item.ResidencyCoursesTeachers {
		item.ResidencyCoursesTeachers[i].ResidencyCourseID = item.ID
	}
	for i := range item.ResidencyCoursesSpecializations {
		item.ResidencyCoursesSpecializations[i].ResidencyCourseID = item.ID
	}
}

func (item *ResidencyCourse) SetFilePath(fileID string) *string {
	if item.Annotation.ID.UUID.String() == fileID {
		item.Annotation.FileSystemPath = uploadHelper.BuildPath(&fileID)
		return &item.Annotation.FileSystemPath
	}
	if item.Plan.ID.UUID.String() == fileID {
		item.Plan.FileSystemPath = uploadHelper.BuildPath(&fileID)
		return &item.Plan.FileSystemPath
	}
	if item.Program.ID.UUID.String() == fileID {
		item.Program.FileSystemPath = uploadHelper.BuildPath(&fileID)
		return &item.Program.FileSystemPath
	}
	if item.Schedule.ID.UUID.String() == fileID {
		item.Schedule.FileSystemPath = uploadHelper.BuildPath(&fileID)
		return &item.Schedule.FileSystemPath
	}
	return nil
}
