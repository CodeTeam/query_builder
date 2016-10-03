package builder

import (
	"bytes"
	"strings"
)

func (query *Query) BuildQuery() string {
	var buffer bytes.Buffer
	buffer.WriteString(query.TypeQuery)
	buffer.WriteString(" ")
	columns := interfaceToString(query.Columns)
	buffer.WriteString(strings.Join(columns[:], ", "))

	buffer.WriteString(" From ")
	buffer.WriteString(query.TableName)

	if len(query.WhereCond) == 0 {
		buffer.WriteString(" Where ")
		for index, element := range query.WhereCond {
			if index == 0 {
				buffer.WriteString(convertValueToString(element.Expression, element.Value))
			} else {
				buffer.WriteString(element.Delimiter)
				buffer.WriteString(convertValueToString(element.Expression, element.Value))
			}
		}
	}
	return buffer.String()
}
