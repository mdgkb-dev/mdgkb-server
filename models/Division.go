package models

import "github.com/google/uuid"

type Division struct {
	ID          uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name        string     `json:"name"`
	Info        string     `json:"info"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	Address     string     `json:"address"`
	Slug        string     `json:"slug"`
	Doctors     []*Doctor  `bun:"rel:has-many" json:"doctors"`
	EntranceId  uuid.UUID  `bun:"type:uuid" json:"entranceId"`
	FloorId     uuid.UUID  `bun:"type:uuid" json:"floorId"`
	Timetable   *Timetable `bun:"rel:belongs-to" json:"timetable"`
	TimetableId uuid.UUID  `bun:"type:uuid" json:"timetableId"`
	Schedule    *Schedule  `bun:"rel:belongs-to" json:"schedule"`
	ScheduleId  uuid.UUID  `bun:"type:uuid" json:"scheduleId"`
}
