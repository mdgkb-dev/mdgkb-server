package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type ResidencyApplicationPointsAchievement struct {
	bun.BaseModel `bun:"residency_applications_points_achievement,alias:residency_applications_points_achievement"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	ResidencyApplication   *ResidencyApplication `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyApplicationID uuid.NullUUID         `bun:"type:uuid" json:"residencyCourseId"`

	PointsAchievement   *PointsAchievement `bun:"rel:belongs-to" json:"pointsAchievement"`
	PointsAchievementID uuid.NullUUID      `bun:"type:uuid" json:"pointsAchievementId"`

	FileInfo   *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`

	Approved bool `json:"approved"`
}

type ResidencyApplicationPointsAchievements []*ResidencyApplicationPointsAchievement

func (items ResidencyApplicationPointsAchievements) SetForeignKeys() {
	for i := range items {
		items[i].FileInfoID = items[i].FileInfo.ID
	}
}

func (items ResidencyApplicationPointsAchievements) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.FileInfo)
	}
	return itemsForGet
}

func (items ResidencyApplicationPointsAchievements) SetFilePath(fileID *string) *string {
	for i := range items {
		filePath := items[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *ResidencyApplicationPointsAchievement) SetFilePath(fileID *string) *string {
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploader.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	return nil
}
