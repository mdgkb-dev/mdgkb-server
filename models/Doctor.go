package models

import "github.com/google/uuid"

type Doctor struct {
	ID             uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Division       *Division        `bun:"rel:belongs-to" json:"division"`
	DivisionId     uuid.UUID        `bun:"type:uuid,nullzero,default:NULL" json:"divisionId,omitempty"`
	Human          *Human           `bun:"rel:belongs-to" json:"human"`
	HumanId        uuid.UUID        `bun:"type:uuid" json:"humanId"`
	Education      string           `json:"education"`
	Schedule       string           `json:"schedule"`
	Position       string           `json:"position"`
	Tags           string           `json:"tags"`
	FileInfo       *FileInfo        `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId     uuid.UUID        `bun:"type:uuid" json:"fileInfoId"`
	DoctorComments []*DoctorComment `bun:"rel:has-many" json:"doctorComments"`
}
