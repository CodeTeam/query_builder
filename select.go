package builder

import (
	"strings"
)

type Query struct {
	TypeQuery     string
	Columns       []interface{}
	TableName     string
	WhereCond     []WhereStruct
	GroupByStruct []interface{}
}

type WhereStruct struct {
	Expression string
	Value      interface{}
	Delimiter  string
}

func Select(columns ...interface{}) *Query {

	return &Query{
		Columns:   columns,
		TypeQuery: "Select",
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
	return query
}

func (query *Query) And(query_str string, value interface{}) *Query {
	query.WhereCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value, Delimiter: " And "})
	return query
}

func (query *Query) Or(query_str string, value interface{}) *Query {
	query.WhereCond = append(query.WhereCond, WhereStruct{Expression: query_str, Value: value, Delimiter: " Or "})
	return query
}

func (query *Query) GroupBy(values ...interface{}) *Query {
	query.GroupByStruct = values
	return query
}
