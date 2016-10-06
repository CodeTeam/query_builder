package builder

import (
	"bytes"
	"strings"
)

type Query struct {
	TypeQuery      string
	Columns        []string
	TableName      string
	WhereCond      []WhereStruct
	HavingCond     []WhereStruct
    ValuesStruct   []string
    ReturningStruct      []string
	GroupByStruct  []interface{}
	DistinctStruct bool
	IsWhere        bool
}

type WhereStruct struct {
	Expression string
	Value      interface{}
	Delimiter  string
}

//BuildQuery - get sql string from expression
func (query *Query) BuildQuery() string {
	var res string
    switch {
    case query.TypeQuery == "select":
        res = buildSelect(query)
    case query.TypeQuery == "update":
        res = buildUpdate(query)
    }

    return res
}

func buildSelect(query *Query) string {
    var buffer bytes.Buffer

	buffer.WriteString("Select")

	if query.DistinctStruct == true {
		buffer.WriteString(" Distinct ")
	} else {
		buffer.WriteString(" ")
	}
	//columns := interfaceToString(query.Columns)
	buffer.WriteString(strings.Join(query.Columns[:], ", "))

	buffer.WriteString(" From ")
	buffer.WriteString(query.TableName)

	if len(query.WhereCond) != 0 {
		buffer.WriteString(buildWhere(query.WhereCond))
	}

	if len(query.GroupByStruct) != 0 {
		buffer.WriteString(" Group By ")
		res := make([]string, len(query.GroupByStruct))

		for index, element := range query.GroupByStruct {
			res[index] = convertValueToString(element)
		}
		buffer.WriteString(strings.Join(res[:], ", "))
	}

	if len(query.HavingCond) != 0 {
		buffer.WriteString(" Having ")
		for index, element := range query.HavingCond {
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
	return buffer.String()
}

func buildUpdate(query *Query) string {
    var buffer bytes.Buffer

	buffer.WriteString("Update ")
    buffer.WriteString(query.TableName)
    buffer.WriteString(" Set (")
    buffer.WriteString(strings.Join(query.Columns[:], ", "))
    buffer.WriteString(") = (")
    buffer.WriteString(strings.Join(query.ValuesStruct[:], ", "))
    buffer.WriteString(")")
    if len(query.WhereCond) != 0 {
		buffer.WriteString(buildWhere(query.WhereCond))
	}
    if len(query.ReturningStruct) != 0 {
        buffer.WriteString(" Returning ")
        buffer.WriteString(strings.Join(query.ReturningStruct[:], ", "))
    }
    
    return buffer.String()
}

func buildWhere(where []WhereStruct) string {
    var buffer bytes.Buffer

    buffer.WriteString(" Where ")
    for index, element := range where {
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
    return buffer.String()
}
