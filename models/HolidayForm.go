package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type HolidayForm struct {
	ID uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Email     string    `json:"email"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`

	Phone            string `json:"phone"`
	Name             string `json:"name"`
	Fio              string `json:"fio"`
	ParentFio        string `json:"parentFio"`
	Surname          string `json:"surname"`
	Patronymic       string `json:"patronymic"`
	ParentName       string `json:"parentName"`
	ParentSurname    string `json:"parentSurname"`
	ParentPatronymic string `json:"parentPatronymic"`
	Representative   string `json:"representative"`

	Dance      string `json:"dance"`
	Song       string `json:"song"`
	Music      string `json:"music"`
	Poem       string `json:"poem"`
	CustomShow string `json:"customShow"`

	NeedingSlice  []string `bun:"-" json:"needing"`
	Needing       string
	CustomNeeding string `json:"customNeeding"`

	Color       string   `json:"color"`
	Hobby       string   `json:"hobby"`
	Happy       string   `json:"happy"`
	PlaceSlice  []string `bun:"-" json:"place"`
	Place       string
	CustomPlace string `json:"customPlace"`
}

func (item *HolidayForm) ToString() {
	item.Place = strings.Join(item.PlaceSlice, ",")
	item.Needing = strings.Join(item.NeedingSlice, ",")
}
