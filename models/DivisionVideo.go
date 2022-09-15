package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DivisionVideo struct {
	bun.BaseModel  `bun:"division_videos,alias:division_videos"`
	ID             uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DivisionID     uuid.NullUUID `bun:"type:uuid" json:"divisionId" `
	Division       *Division     `bun:"rel:belongs-to" json:"division" `
	YouTubeVideoID string        `json:"youTubeVideoId"`
}

type DivisionVideos []*DivisionVideo

func (items DivisionVideos) GetYouTubeVideoIDs() []string {
	ids := make([]string, 0)
	for _, item := range items {
		if item.YouTubeVideoID != "" {
			ids = append(ids, item.YouTubeVideoID)
		}
	}
	return ids
}
