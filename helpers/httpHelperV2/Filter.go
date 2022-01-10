package httpHelper

import (
	"fmt"

	"github.com/uptrace/bun"
)

//
// CreateOrder func
func CreateOrder(query *bun.SelectQuery, sortModels SortModels, defaultSort []string) {
	if len(sortModels) != 0 {
		for _, sort := range sortModels {
			query = query.Order(fmt.Sprintf("%s %s", sort.GetTableAndCol(), *sort.Order))
		}
		return
	}
	for _, sort := range defaultSort {
		query = query.Order(sort)
	}
}

// CreateFilter func
func CreateFilter(query *bun.SelectQuery, filterModels FilterModels) {
	if len(filterModels) == 0 {
		return
	}
	for _, filter := range filterModels {
		switch *filter.Type {
		case SetType:
			if len(filter.Set) == 0 {
				break
			}
			constructWhereIn(query, filter)
		case DateType:
			filter.DatesToString()
			constructWhere(query, filter)
		case StringType:
			if filter.IsLike() {
				filter.LikeToString()
			}
			constructWhere(query, filter)
		//case "number":
		//	tbl = constructNumberWhere(tbl, field, filter)
		//case "text":
		//	if filterOperator == "" {
		//		tbl = constructTextWhere(tbl, field, filterOperator, filter)
		//	} else {
		//		tbl = constructTextWhere(tbl, field, filterOperator, filter.Condition1.Filter, filter.Condition2.Filter)
		//	}
		default:
			//log.Println("unknown number filterType: " + *filter.FilterType)
			return
		}
	}
	return
}

func constructWhere(query *bun.SelectQuery, filter *FilterModel) {
	q := ""
	if filter.IsUnary() {
		q = fmt.Sprintf("%s %s '%s'", filter.GetTableAndCol(), *filter.Operator, filter.Value1)
	}
	if filter.IsBetween() {
		q = fmt.Sprintf("%s %s '%s' and '%s'", filter.GetTableAndCol(), *filter.Operator, filter.Value1, filter.Value2)
	}
	query = query.Where(q)
}

func constructWhereIn(query *bun.SelectQuery, filter *FilterModel) {
	if *filter.JoinTable == "" {
		query = query.Where(fmt.Sprintf("%s %s (?)", filter.GetTableAndCol(), *filter.Operator), bun.In(filter.Set))
		return
	}
	q := fmt.Sprintf("EXISTS (SELECT NULL from patient_diagnosis where %s and %s in (?))", filter.GetJoinCondition(), filter.GetTableAndCol())
	query = query.Where(q, bun.In(filter.Set))
}

//
//func constructTextWhere(tbl *bun.SelectQuery, field string, operator string, options ...models.Filter) *bun.SelectQuery {
//	operators := map[string]string{
//		"equals":      "%s = ?",
//		"notEqual":    "%s <> ?",
//		"contains":    "%s LIKE ?",
//		"notContains": "%s NOT LIKE ?",
//		"startsWith":  "%s LIKE ?",
//		"endsWith":    "%s LIKE ?",
//	}
//	if operator == "" {
//		tbl = constructQuery(tbl, operators[*options[0].Type], "", field, operator, likeMix(*options[0].Type, fmt.Sprintf("%v", *options[0].Filter)))
//	} else {
//		tbl = constructQuery(tbl, operators[*options[0].Type], operators[*options[1].Type], field, operator, likeMix(*options[0].Type, (*options[0].Filter).(string)), likeMix(*options[1].Type, (*options[1].Filter).(string)))
//	}
//	return tbl
//}
//
//func likeMix(typeOperator, filter string) (result string) {
//	likeOperators := map[string]string{
//		"contains":    "%%%s%%",
//		"notContains": "%%%s%%",
//		"startsWith":  "%s%%",
//		"endsWith":    "%%%s",
//	}
//	likePhrase, ok := likeOperators[typeOperator]
//	if ok {
//		return fmt.Sprintf(likePhrase, filter)
//	}
//	return filter
//}
//
//func getTimeString(date string) string {
//	resultDate := time.Now()
//	resultDate, err := time.Parse("2006-01-02 15:04:05", date)
//	if err != nil {
//		resultDate = time.Date(resultDate.Year(), resultDate.Month(), resultDate.Day(), 0, 0, 0, 0, time.UTC)
//	}
//	location, _ := time.LoadLocation("Europe/Moscow")
//	result := resultDate.In(location).Format("2006-01-02")
//	return result
//}
//
