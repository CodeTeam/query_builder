package main

import (
	"bytes"
	"fmt"

	qb "github.com/CodeTeam/query_builder"
)

func main() {
	var buffer bytes.Buffer
	fmt.Println(buffer.String())

	fmt.Println("test")

	subquery := qb.Select("f1").From("table2")
	subquery1 := qb.Select("field1", "field2", "field3", "field4").Distinct().From("table2")

	b := qb.Select("field1", "field2", "field3", "field4").
		From("table 1").
		Where("field1 = ?", 1.0).
		Or("field3 = ?", "sfsdfds").
		And("field4 IN (?)", []int{1, 2, 3, 4, 5}).
		And("field2 IN (?)", subquery.BuildQuery())

	fmt.Println(b.BuildQuery())
	fmt.Println(b.DistinctStruct)

	b1 := qb.Select("field1", "field2", "field3", "field4").
		FromSubquery(subquery1.BuildQuery()).
		Where("field1 = ?", 1).
		Or("field3 = ?", "sfsdfds").
		And("field4 IN (?)", []float64{1.0, 2.0, 3.0, 4.0, 5.0}).
		Or("field2 IN (?)", subquery1.BuildQuery()).
		GroupBy("field1").GroupBy("field2").GroupBy("field3").
		Having("field1 = ?", 1)

	fmt.Println(b1.BuildQuery())

}
