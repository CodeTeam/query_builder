package main

import (
    "fmt"
    "strings"
)

type Query struct {
    TypeQuery string
    Columns []interface{}
    TableName string
}

func main() {
   fmt.Println("test")
   
   b := Select("sadasd", "asdasdsad").From("table 1")
   fmt.Println(b)
   fmt.Println(b.Columns)
   fmt.Println(b.BuildQuery())
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
    sql := query.TypeQuery + " "
    columns := interfaceToString(query.Columns)
    sql = sql + strings.Join(columns[:], ", ")

    sql = sql + " From " + query.TableName
    return sql
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
