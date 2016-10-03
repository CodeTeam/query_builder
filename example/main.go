package main

import (
	"bytes"
	"fmt"

	qb "github.com/CodeTeam/query_builder"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("sdfsdfs")
	buffer.WriteString(" 111")
	fmt.Println(buffer.String())

	fmt.Println("test")

	subquery := qb.Select("f1").From("table2")
	subquery1 := qb.Select("field1", "field2", "field3", "field4").From("table2")

	b := qb.Select("field1", "field2", "field3", "field4").
		From("table 1").
		Where("field1 = ?", 1).
		Or("field3 = ?", "sfsdfds").
		And("field4 IN (?)", []int{1, 2, 3, 4, 5}).
		And("field2 IN (?)", subquery.BuildQuery())

	fmt.Println(b.BuildQuery())

	b1 := qb.Select("field1", "field2", "field3", "field4").
		FromSubquery(subquery1.BuildQuery()).
		Where("field1 = ?", 1).
		Or("field3 = ?", "sfsdfds").
		And("field4 IN (?)", []int{1, 2, 3, 4, 5}).
		And("field2 IN (?)", subquery.BuildQuery())

	fmt.Println(b1.BuildQuery())

}

/*func interfaceToString(input_int []interface{}) []string {
    length := len(input_int)
    columns := make([]string, length)
    for index, element := range input_int {
        columns[index] = element.(string)
    }
    return columns
}

func (query *Query) BuildQuery() string {
    var buffer bytes.Buffer
    buffer.WriteString(query.TypeQuery)
    buffer.WriteString(" ")
    columns := interfaceToString(query.Columns)
    buffer.WriteString(strings.Join(columns[:], ", "))

    buffer.WriteString(" From ")
    buffer.WriteString(query.TableName)

    buffer.WriteString(" Where ")
    for index, element := range query.WhereCond {
        if index == 0 {
            buffer.WriteString(convertValueToString(element.Expression, element.Value))
        } else {
            buffer.WriteString(element.Delimiter)
            buffer.WriteString(convertValueToString(element.Expression, element.Value))
        }
    }
    return buffer.String()
}

func convertValueToString(expr string, value interface{}) string {
    var result string
    switch value := value.(type) {
        case int:
            result = strings.Replace(expr, "?", strconv.Itoa(int(value)), -1)
        case string:
            result = strings.Replace(expr, "?", value, -1)
        case []int:
            var res bytes.Buffer
            for index, el := range value {
                res.WriteString(strconv.Itoa(int(el)))
                if index != len(value) -1 {
                    res.WriteString(", ")
                }
            }
            result = strings.Replace(expr, "?", res.String(), -1)

    }
    return result
}

func Select(columns ...interface{}) *Query {

	return &Query{
		Columns:      columns,
        TypeQuery: "Select",
	}
}

func (query *Query) From(table string) *Query {
    query.TableName = strings.Replace(table, " ", "", -1)
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
}*/
