package models

import (
	"time"

	"github.com/google/uuid"
)

type HolidayForm struct {
	ID uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Email     string    `json:"email"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`

	Phone            string `json:"phone"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Patronymic       string `json:"patronymic"`
	ParentName       string `json:"parentName"`
	ParentSurname    string `json:"parentSurname"`
	ParentPatronymic string `json:"parentPatronymic"`
	Representative   string `json:"representative"`

	Dance      string `json:"dance"`
	Song       string `json:"song"`
	Music      string `json:"music"`
	CustomShow string `json:"customShow"`

	Needing string `json:"needing"`

	Color string `json:"color"`
	Hobby string `json:"hobby"`
	Happy string `json:"happy"`
	Place string `json:"place"`
}
