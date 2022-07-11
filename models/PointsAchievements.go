package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PointsAchievement struct {
	bun.BaseModel `bun:"points_achievements,alias:points_achievements"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Name   string `json:"name"`
	Points int    `json:"points"`
	Code   string `json:"code"`
	Order  uint8  `bun:"points_achievements_order" json:"order"`
}

type PointsAchievements []*PointsAchievement
