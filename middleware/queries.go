package middleware

import (
	"context"
	"net/http"

	"github.com/pro-assistance/pro-assister/sqlHelper"
)

type Query struct {
	QID  string                 `json:"qid"`
	FTSP *sqlHelper.QueryFilter `json:"ftsp"`
}

var queriesMap = make(map[string]*sqlHelper.QueryFilter)

func (item Query) Inject(r *http.Request, qid string) error {
	ftsp := getFTSP(qid)
	if ftsp == nil {
		// setFTSP(nil)
	}
	*r = *r.WithContext(context.WithValue(r.Context(), "ftsp", ftsp))
	return nil
}

func getFTSP(qid string) *sqlHelper.QueryFilter {
	filter, ok := queriesMap[qid]
	if !ok {
		return nil
	}
	return filter
}

func setFTSP(query Query) {
	queriesMap[query.QID] = query.FTSP
}
