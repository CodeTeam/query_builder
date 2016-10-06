package builder

import (
	"strings"
)

// Select - func add Select to sql
func Select(columns ...interface{}) *Query {

	return &Query{
		Columns:   interfaceToString(columns),
		TypeQuery: "select",
	}
}

// From - func add From to sql
func (query *Query) From(table string) *Query {
	query.TableName = strings.Replace(table, " ", "", -1)
	return query
}

// FromSubquery - func add FromSubquery to sql
func (query *Query) FromSubquery(table string) *Query {
	query.TableName = "(" + table + ")"
	return query
}

// Where - func add Where to sql
func (query *Query) Where(queryStr string, value interface{}) *Query {
	query.WhereCond = append(query.WhereCond, WhereStruct{Expression: queryStr, Value: value})
	query.IsWhere = true
	return query
}

// And - func add And to sql
func (query *Query) And(queryStr string, value interface{}) *Query {
	if query.IsWhere == true {
		query.WhereCond = append(query.WhereCond, WhereStruct{Expression: queryStr, Value: value, Delimiter: " And "})
	} else {
		query.HavingCond = append(query.WhereCond, WhereStruct{Expression: queryStr, Value: value, Delimiter: " And "})
	}
	return query
}

// Or - func add Or to sql
func (query *Query) Or(queryStr string, value interface{}) *Query {
	if query.IsWhere == true {
		query.WhereCond = append(query.WhereCond, WhereStruct{Expression: queryStr, Value: value, Delimiter: " Or "})
	} else {
		query.HavingCond = append(query.WhereCond, WhereStruct{Expression: queryStr, Value: value, Delimiter: " Or "})
	}
	return query
}

// GroupBy - func add GroupBy to sql
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
