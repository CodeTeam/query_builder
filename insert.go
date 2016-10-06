package builder

import (
	"strings"
)

func Insert(table string) *Query {
	return &Query{
		TableName: table,
		TypeQuery: "insert",
	}
}

func (query *Query) Record(values ...interface{}) *Query {
    record := interfaceToString(values)
    record_str := strings.Join(record, ", ")
    query.RecordsSrtuct = append(query.RecordsSrtuct, "(" + record_str + ")")
	return query
}
