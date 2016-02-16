package df

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// column represents a column inside a DataFrame
type column struct {
	cells    cells
	colType  string
	colName  string
	numChars int
}

type columns []column

// newCol is the constructor for a new Column with the given colName and elements
func newCol(colName string, elements cells) (*column, error) {
	col := column{
		colName: colName,
	}
	col, err := col.append(elements...)
	if err != nil {
		return nil, err
	}

	return &col, nil
}

// Implementing the Stringer interface for Column
func (col column) String() string {
	strArray := []string{}

	for i := 0; i < len(col.cells); i++ {
		strArray = append(strArray, col.cells[i].String())
	}

	return fmt.Sprintln(
		col.colName,
		"(", col.colType, "):\n",
		strings.Join(strArray, "\n "),
	)
}

func parseColumn(col column, t string) (*column, error) {
	switch t {
	case "string":
		newcells := Strings(col.cells)
		newcol, err := newCol(col.colName, newcells)
		return newcol, err
	case "int":
		newcells := Ints(col.cells)
		newcol, err := newCol(col.colName, newcells)
		return newcol, err
	case "float":
		newcells := Floats(col.cells)
		newcol, err := newCol(col.colName, newcells)
		return newcol, err
	case "bool":
		newcells := Bools(col.cells)
		newcol, err := newCol(col.colName, newcells)
		return newcol, err
	}
	return nil, errors.New("Can't parse the given type")
}

func (col *column) recountNumChars() {
	numChars := len(col.colName)
	for _, cell := range col.cells {
		cellStr := cell.String()
		if len(cellStr) > numChars {
			numChars = len(cellStr)
		}
	}

	col.numChars = numChars
}

// Append will add a value or values to a column
func (col column) append(values ...cell) (column, error) {
	if len(values) == 0 {
		col.recountNumChars()
		return col, nil
	}

	for _, v := range values {
		t := reflect.TypeOf(v).String()
		if col.colType == "" {
			col.colType = t
		} else {
			if t != col.colType {
				return col, errors.New("Can't have elements of different type on the same column")
			}
		}

		col.cells = append(col.cells, v)
	}

	col.recountNumChars()

	return col, nil
}
