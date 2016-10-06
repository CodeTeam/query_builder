package builder


func Update(table string) *Query {
	return &Query{
		TableName: table,
		TypeQuery: "update",
	}
}

func (query *Query) Fields(values ...interface{}) *Query {
	query.Columns = interfaceToString(values)
	return query
}

func (query *Query) Values(values ...interface{}) *Query {
	query.ValuesStruct = interfaceToString(values)
	return query
}

func (query *Query) Returning(values ...interface{}) *Query {
	query.ReturningStruct = interfaceToString(values)
	return query
}
