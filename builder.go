package builder

import (
	"bytes"
	"strings"
)

//BuildQuery - get sql string from expression
func (query *Query) BuildQuery() string {
	var buffer bytes.Buffer
	buffer.WriteString(query.TypeQuery)
	buffer.WriteString(" ")
	columns := interfaceToString(query.Columns)
	buffer.WriteString(strings.Join(columns[:], ", "))

	buffer.WriteString(" From ")
	buffer.WriteString(query.TableName)

	if len(query.WhereCond) != 0 {
		buffer.WriteString(" Where ")
		for index, element := range query.WhereCond {
			if index == 0 {
				buffer.WriteString(
					strings.Replace(element.Expression, "?", convertValueToString(element.Value), -1),
				)
			} else {
				buffer.WriteString(element.Delimiter)
				buffer.WriteString(
					strings.Replace(element.Expression, "?", convertValueToString(element.Value), -1),
				)
			}
		}
	}

	if len(query.GroupByStruct) != 0 {
		buffer.WriteString(" Group By ")
		for index, element := range query.GroupByStruct {
			buffer.WriteString(
				convertValueToString(element),
			)
		}
	}
	buffer.WriteString(";")
	return buffer.String()
}
