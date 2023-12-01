package models

import (
	"fmt"

	"github.com/ip2location/ip2location-go/v9"
)

type GeoIP struct {
	// The right side is the name of the JSON variable
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
}

func (item *GeoIP) GetByIP(ipStr string) (string, string) {
	db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB3.BIN")
	if err != nil {
		fmt.Println(err)
	}
	results, err := db.Get_all(ipStr)
	if err != nil {
		fmt.Println(err)
	}

	db.Close()

	return results.Country_long, results.City
}
