package middleware

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/sqlHelper"
)

type FTSPStore struct {
	store map[string]sqlHelper.FTSP
}

var ftspStore = FTSPStore{store: make(map[string]sqlHelper.FTSP)}

func (item FTSPStore) SetFTSP(query *sqlHelper.FTSPQuery) {
	id := uuid.NewString()
	query.FTSP.ID = id
	query.QID = id
	item.store[id] = query.FTSP
}

func (item FTSPStore) GetFTSP(qid string) (sqlHelper.FTSP, bool) {
	fmt.Println(item.store)
	ftsp, ok := item.store[qid]
	return ftsp, ok
}

func (item FTSPStore) GetOrCreateFTSP(query *sqlHelper.FTSPQuery) (sqlHelper.FTSP, bool) {
	if query.QID == "" {
		item.SetFTSP(query)
		return query.FTSP, true
	}
	return item.GetFTSP(query.QID)
}
