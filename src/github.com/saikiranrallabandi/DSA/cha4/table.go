package main
import (
	"fmt"
)

// Table class
type Table struct {
	Rows []Row
	Name string
	ColumnNames []string
}

// Row class
type Row struct {
	Columns []Column
	Id int
}

// Column class
type Column struct {
	Id int
	Value string
}

//Print Table
func printTable(table Table)  {
	var rows []Row = table.Rows
	fmt.Println(table.Name)
	for _,row := range rows {
		var columns []Column = row.Columns
		for i, column := range columns {
			fmt.Println(table.ColumnNames[i],column.Id, column.Value);
		}
	}

}
func main2() {
	var table Table = Table{}
	table.Name = "Customer"
	table.ColumnNames = []string{"Id", "Name", "SSN"}

	var rows []Row = make([]Row,2)
	rows[0] = Row{}
	var columns1 []Column = make([]Column, 3)
	columns1[0] = Column{1,"323"}
	rows[0].Columns = columns1
	rows[1] = Row{}
	table.Rows = rows
	fmt.Println(table)
	printTable(table)
}