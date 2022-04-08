package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyCourse struct {
	bun.BaseModel                            `bun:"residency_courses,select:residency_courses_view,alias:residency_courses_view"`
	ID                                       uuid.NullUUID                   `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Description                              string                          `json:"description"`
	EducationForm                            string                          `json:"educationForm"`
	Years                                    int                             `json:"years"`
	Slug                                     string                          `bun:",scanonly" json:"slug"`
	Listeners                                int                             `json:"listeners"`
	ResidencyCoursesSpecializations          ResidencyCoursesSpecializations `bun:"rel:has-many" json:"residencyCoursesSpecializations"`
	ResidencyCoursesSpecializationsForDelete []uuid.UUID                     `bun:"-" json:"residencyCoursesSpecializationsForDelete"`
	ResidencyCoursesTeachers                 ResidencyCoursesTeachers        `bun:"rel:has-many" json:"residencyCoursesTeachers"`
	ResidencyCoursesTeachersForDelete        []uuid.UUID                     `bun:"-" json:"residencyCoursesForDelete"`

	//QuestionsFile   *FileInfo     `bun:"rel:belongs-to" json:"questionsFile"`
	//QuestionsFileID uuid.NullUUID `bun:"type:uuid" json:"questionsFileId"`
	//
	//ProgramFile   *FileInfo     `bun:"rel:belongs-to" json:"programFile"`
	//ProgramFileID uuid.NullUUID `bun:"type:uuid" json:"programFileId"`
	//
	//Calendar   *FileInfo     `bun:"rel:belongs-to" json:"calendar"`
	//CalendarID uuid.NullUUID `bun:"type:uuid" json:"calendarId"`

	//ResidencyCoursePlans          ResidencyCoursePlans `bun:"rel:has-many" json:"residencyCoursePlans"`
	//ResidencyCoursePlansForDelete []uuid.UUID             `bun:"-" json:"residencyCoursePlansForDelete"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`

	//DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	//DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentTypeId"`
}

type ResidencyCourses []*ResidencyCourse

func (item *ResidencyCourse) SetForeignKeys() {
	//item.QuestionsFileID = item.QuestionsFile.ID
	//item.ProgramFileID = item.ProgramFile.ID
	//item.CalendarID = item.Calendar.ID
	item.FormPatternID = item.FormPattern.ID
	//item.DocumentTypeID = item.DocumentType.ID
}

func (item *ResidencyCourse) SetIdForChildren() {
	for i := range item.ResidencyCoursesTeachers {
		item.ResidencyCoursesTeachers[i].ResidencyCourseID = item.ID
	}
	for i := range item.ResidencyCoursesSpecializations {
		item.ResidencyCoursesSpecializations[i].ResidencyCourseID = item.ID
	}
	//for i := range item.ResidencyCoursePlans {
	//	item.ResidencyCoursePlans[i].ResidencyCourseID = item.ID
	//}
}

func (item *ResidencyCourse) SetFilePath(fileID string) *string {
	//if item.QuestionsFile.ID.UUID.String() == fileID {
	//	item.QuestionsFile.FileSystemPath = uploadHelper.BuildPath(&fileID)
	//	return &item.QuestionsFile.FileSystemPath
	//}
	//if item.Calendar.ID.UUID.String() == fileID {
	//	item.Calendar.FileSystemPath = uploadHelper.BuildPath(&fileID)
	//	return &item.Calendar.FileSystemPath
	//}
	//if item.ProgramFile.ID.UUID.String() == fileID {
	//	item.ProgramFile.FileSystemPath = uploadHelper.BuildPath(&fileID)
	//	return &item.ProgramFile.FileSystemPath
	//}
	//filePath := item.ResidencyCoursePlans.SetFilePath(fileID)
	//if filePath != nil {
	//	return filePath
	//}
	return nil
}
