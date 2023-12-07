package exportmodels

import (
	"encoding/json"
	"errors"
)

type NewsView struct {
	IDPool []string           `json:"ids"`
	Type   NewsViewTypeExport `json:"type"`
}

type NewsViewTypeExport string

const (
	NewsViewTypeExportCities NewsViewTypeExport = "cities"
	NewsViewTypeExportDates  NewsViewTypeExport = "dates"
	NewsViewTypeExportHours  NewsViewTypeExport = "hours"
)

func (item *NewsView) GetColExpr() string {
	if item.Type == NewsViewTypeExportCities {
		return "city as label, 100.0 * count(*) / sum(count(*)) over() as value"
	}

	if item.Type == NewsViewTypeExportDates {
		return "nv.created_at::date as label, count(nv.created_at::date) as value"
	}
	return ""
}

func (item *NewsView) GetGroupExpr() string {
	if item.Type == NewsViewTypeExportCities {
		return "label"
	}

	if item.Type == NewsViewTypeExportDates {
		return "label"
	}
	return ""
}

const newsViewKey = "newsView"

func (item *NewsView) ParseExportOptions(options map[string]map[string]interface{}) error {
	opt, ok := options[newsViewKey]
	if !ok {
		return errors.New("not find parametrs")
	}
	jsonbody, err := json.Marshal(opt[newsViewKey])
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonbody, &item)
	return err
}
