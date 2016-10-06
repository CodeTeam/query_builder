package builder

import (
	"strings"
)

func Select(columns ...interface{}) *Query {

	return &Query{
		Columns:   interfaceToString(columns),
		TypeQuery: "select",
	}
}

func (query *Query) From(table string) *Query {
	query.TableName = strings.Replace(table, " ", "", -1)
	return query
}

func (query *Query) FromSubquery(table string) *Query {
	query.TableName = "(" + table + ")"
	return query
}

func (query *Query) Where(query_str string, value interface{}) *Query {
	query.WhereCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value})
	query.IsWhere = true
	return query
}

func (query *Query) And(query_str string, value interface{}) *Query {
	if query.IsWhere == true {
		query.WhereCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value, Delimiter: " And "})
	} else {
		query.HavingCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value, Delimiter: " And "})
	}
	return query
}

func (query *Query) Or(query_str string, value interface{}) *Query {
	if query.IsWhere == true {
		query.WhereCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value, Delimiter: " Or "})
	} else {
		query.HavingCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value, Delimiter: " Or "})
	}
	return query
}

func (query *Query) GroupBy(value interface{}) *Query {
	query.GroupByStruct = append(query.GroupByStruct, value)
	return query
}

// Distinct - add Distinct to sql queru
func (query *Query) Distinct() *Query {
	query.DistinctStruct = true
	return query
}

// Having - having sql expression
func (query *Query) Having(queryStr string, value interface{}) *Query {
	query.HavingCond = append(query.HavingCond, WhereStruct{Expression: queryStr, Value: value})
	query.IsWhere = false
	return query
}
