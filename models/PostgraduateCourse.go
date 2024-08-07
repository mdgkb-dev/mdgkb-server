package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type PostgraduateCourse struct {
	bun.BaseModel                               `bun:"postgraduate_courses,select:postgraduate_courses_view,alias:postgraduate_courses_view"`
	ID                                          uuid.NullUUID                      `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Description                                 string                             `json:"description"`
	EducationForm                               string                             `json:"educationForm"`
	Years                                       int                                `json:"years"`
	Slug                                        string                             `bun:",scanonly" json:"slug"`
	Cost                                        int                                `json:"cost"`
	PostgraduateCoursesSpecializations          PostgraduateCoursesSpecializations `bun:"rel:has-many" json:"postgraduateCoursesSpecializations"`
	PostgraduateCoursesSpecializationsForDelete []uuid.UUID                        `bun:"-" json:"postgraduateCoursesSpecializationsForDelete"`
	PostgraduateCoursesTeachers                 PostgraduateCoursesTeachers        `bun:"rel:has-many" json:"postgraduateCoursesTeachers"`
	PostgraduateCoursesTeachersForDelete        []uuid.UUID                        `bun:"-" json:"postgraduateCoursesForDelete"`
	PostgraduateCoursesDates                    PostgraduateCoursesDates           `bun:"rel:has-many" json:"postgraduateCoursesDates"`
	PostgraduateCoursesDatesForDelete           []uuid.UUID                        `bun:"-" json:"postgraduateCoursesDatesForDelete"`

	Annotation   *FileInfo     `bun:"rel:belongs-to" json:"annotation"`
	AnnotationID uuid.NullUUID `bun:"type:uuid" json:"annotationId"`

	QuestionsFile   *FileInfo     `bun:"rel:belongs-to" json:"questionsFile"`
	QuestionsFileID uuid.NullUUID `bun:"type:uuid" json:"questionsFileId"`

	ProgramFile   *FileInfo     `bun:"rel:belongs-to" json:"programFile"`
	ProgramFileID uuid.NullUUID `bun:"type:uuid" json:"programFileId"`

	Calendar   *FileInfo     `bun:"rel:belongs-to" json:"calendar"`
	CalendarID uuid.NullUUID `bun:"type:uuid" json:"calendarId"`

	PostgraduateCoursePlans          PostgraduateCoursePlans `bun:"rel:has-many" json:"postgraduateCoursePlans"`
	PostgraduateCoursePlansForDelete []uuid.UUID             `bun:"-" json:"postgraduateCoursePlansForDelete"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`

	DocumentType   *PageSection  `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentTypeId"`
}

type PostgraduateCourses []*PostgraduateCourse

func (item *PostgraduateCourse) SetForeignKeys() {
	item.AnnotationID = item.Annotation.ID
	item.QuestionsFileID = item.QuestionsFile.ID
	item.ProgramFileID = item.ProgramFile.ID
	item.CalendarID = item.Calendar.ID
	item.FormPatternID = item.FormPattern.ID
	item.DocumentTypeID = item.DocumentType.ID
}

func (item *PostgraduateCourse) SetIDForChildren() {
	for i := range item.PostgraduateCoursesTeachers {
		item.PostgraduateCoursesTeachers[i].PostgraduateCourseID = item.ID
	}
	for i := range item.PostgraduateCoursesSpecializations {
		item.PostgraduateCoursesSpecializations[i].PostgraduateCourseID = item.ID
	}
	for i := range item.PostgraduateCoursesDates {
		item.PostgraduateCoursesDates[i].PostgraduateCourseID = item.ID
	}
	for i := range item.PostgraduateCoursePlans {
		item.PostgraduateCoursePlans[i].PostgraduateCourseID = item.ID
	}
}

func (item *PostgraduateCourse) SetFilePath(fileID string) *string {
	if item.QuestionsFile.ID.UUID.String() == fileID {
		item.QuestionsFile.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.QuestionsFile.FileSystemPath
	}
	if item.Calendar.ID.UUID.String() == fileID {
		item.Calendar.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Calendar.FileSystemPath
	}
	if item.ProgramFile.ID.UUID.String() == fileID {
		item.ProgramFile.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.ProgramFile.FileSystemPath
	}
	if item.Annotation.ID.UUID.String() == fileID {
		item.Annotation.FileSystemPath = uploader.BuildPath(&fileID)
		return &item.Annotation.FileSystemPath
	}
	filePath := item.PostgraduateCoursePlans.SetFilePath(fileID)
	if filePath != nil {
		return filePath
	}
	return nil
}
