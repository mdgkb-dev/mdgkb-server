package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/pro-assistance/pro-assister/uploadHelper"
)

type DpoApplication struct {
	bun.BaseModel `bun:"dpo_applications,alias:dpo_applications"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	Application   *FileInfo     `bun:"rel:belongs-to" json:"application"`
	ApplicationID uuid.NullUUID `json:"applicationId"`

	OrganizationApplication   *FileInfo     `bun:"rel:belongs-to" json:"organizationApplication"`
	OrganizationApplicationID uuid.NullUUID `bun:"type:uuid" json:"organizationApplicationId"`

	DpoCourse   *DpoCourse    `bun:"rel:belongs-to" json:"dpoCourse"`
	DpoCourseID uuid.NullUUID `bun:"type:uuid" json:"dpoCourseId"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`
}

type DpoApplications []*DpoApplication

func (item *DpoApplication) SetForeignKeys() {
	item.UserID = item.User.ID
	item.DpoCourseID = item.DpoCourse.ID
	item.ApplicationID = item.Application.ID
	item.OrganizationApplicationID = item.OrganizationApplication.ID
}

func (item *DpoApplication) SetFilePath(fileID *string) *string {
	if item.Application.ID.UUID.String() == *fileID {
		item.Application.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Application.FileSystemPath
	}
	if item.OrganizationApplication.ID.UUID.String() == *fileID {
		item.OrganizationApplication.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.OrganizationApplication.FileSystemPath
	}
	return nil
}

func (item *DpoApplication) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.Application, item.OrganizationApplication)
	return items
}
