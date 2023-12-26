package models

import "github.com/pro-assistance/pro-assister/sqlHelper"

type FTSPQuery struct {
	QID  string         `json:"qid"`
	FTSP sqlHelper.FTSP `json:"ftsp"`
}
type FTSPAnswer struct {
	Data interface{} `json:"data"`
}
