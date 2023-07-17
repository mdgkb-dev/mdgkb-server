package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AddressInfo struct {
	bun.BaseModel `bun:"address_infos,alias:address_infos"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Region        string    `json:"region"`
	City          string    `json:"city"`
	Street        string    `json:"street"`
	Building      string    `json:"building"`
	Flat          string    `json:"flat"`
	Zip           int       `json:"zip"`

	RegionID   string `json:"regionId"`
	CityID     string `json:"cityId"`
	StreetID   string `json:"streetId"`
	BuildingID string `bun:"b_id" json:"buildingId"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"cii,type:uuid" json:"contactInfoId"`
}

type AddressInfos []*AddressInfo

func (item *AddressInfo) GetFullAddress() string {
	return fmt.Sprintf("%d, %s, %s, %s, %s", item.Zip, item.Region, item.City, item.Street, item.Building)
}
