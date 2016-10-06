# Query Builder - 
библиотека для упрощения написания sql запросов, ориентированная на запросы
для postgreqsl

1. Select - используется для написания select запросов
*Пример* 
b := qb.Select("field1", "field2", "field3", "field4").
        From("table 1").
		Where("field1 = ?", 1.0).
		Or("field3 = ?", "sfsdfds").
		And("field4 IN (?)", []int{1, 2, 3, 4, 5}).
		And("field2 IN (?)", subquery.BuildQuery())

fmt.Println(b.BuildQuery())
*Результат*
Select field1, field2, field3, field4 From table1 Where field1 = 1.000000 Or field3 = sfsdfds And field4 IN (1, 2, 3, 4, 5) And field2 IN (Select f1 From table2)

2. Update
*Пример* 
u1 := qb.Update("table2").Fields("field1", "field2", "field3", "field4").
		Values("st1", "st2", 10, 45.89).
		Where("field1 = ?", 1).
		Or("field3 = ?", "sfsdfds").
		Returning("field1", "field2")
fmt.Println(u1.BuildQuery())
*Результат*
Update table2 Set (field1, field2, field3, field4) = (st1, st2, 10, 45.890000) Where field1 = 1 Or field3 = sfsdfds Returning field1, field2

3. Insert
*Пример* 
i1 := qb.Insert("table3").
		Fields("field1", "field2", "field3", "field4").
		Record("st1", "st2", 10, 45.89).
		Record("st1", "st2", 10, 45.89).
		Returning("field1")
fmt.Println(i1.BuildQuery())
*Результат*
Insert Into table3 (field1, field2, field3, field4 ) Values (st1, st2, 10, 45.890000), (st1, st2, 10, 45.890000) Returning field1

4. Delete
*Пример* 

*Результат*
