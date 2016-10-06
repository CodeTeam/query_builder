package builder


func Delete(table string) *Query {
	return &Query{
		TableName: table,
		TypeQuery: "delete",
	}
}
