package models

import "fmt"

type KladrAPI struct {
	RegionID   string `json:"regionId"`
	DistrictID string `json:"districtId"`
	CityID     string `json:"cityId"`
	StreetID   string `json:"streetId"`
	BuildingID string `json:"buildingId"`
	Query      string `json:"query"`
}

const freeAPI = "https://kladr-api.ru/api.php"

type kladrContentType string

const (
	region   kladrContentType = "region"
	city     kladrContentType = "city"
	street   kladrContentType = "street"
	building kladrContentType = "building"
)

func (item *KladrAPI) GetURL() string {
	limit := 10
	contentType := item.getContentType()
	codes := item.getCodes()
	query := fmt.Sprintf("?query=%s&contentType=%s&limit=%d&%s", item.Query, contentType, limit, codes)
	return freeAPI + query
}

func (item *KladrAPI) getCodes() string {
	codes := ""
	if item.StreetID != "" {
		codes = codes + "streetId=" + item.StreetID + "&"
	}
	if item.CityID != "" {
		codes = codes + "cityId=" + item.CityID + "&"
	}
	if item.RegionID != "" {
		codes = codes + "regionId=" + item.RegionID + "&"
	}
	return codes
}

func (item *KladrAPI) getContentType() kladrContentType {
	if item.StreetID != "" {
		return building
	}
	if item.CityID != "" {
		return street
	}
	if item.RegionID != "" {
		return city
	}
	return region
}
