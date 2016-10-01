package main

import (
    "fmt"
    "strings"
    "bytes"
    "strconv"
)

type Query struct {
    TypeQuery string
    Columns []interface{}
    TableName string
    WhereCond []WhereStruct
}

type WhereStruct struct {
    Expression string
    Value interface{}
}

func main() {
    var buffer bytes.Buffer
    buffer.WriteString("sdfsdfs")
    buffer.WriteString(" 111")
    fmt.Println(buffer.String())

    fmt.Println("test")
   
    b := Select("field1", "field2").From("table 1").Where("field1 = ?", 1)
    fmt.Println(b)
    fmt.Println(b.Columns)
    fmt.Println(b.BuildQuery())
    fmt.Println(b.WhereCond)
}

func interfaceToString(input_int []interface{}) []string {
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
    for _, element := range query.WhereCond {
        buffer.WriteString(convertValueToString(element.Expression, element.Value))
    }
    return buffer.String()
}

func convertValueToString(expr string, value interface{}) string {
    var result string
    switch value := value.(type) {
        case int:          
            result = strings.Replace(expr, "?", strconv.Itoa(int(value)), -1)

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
