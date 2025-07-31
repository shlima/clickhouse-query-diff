package parser

import "strings"

type Columns []Column

type Column struct {
	Alias string
	SQL   string
}

func (c Columns) String() string {
	out := new(strings.Builder)
	for i := range c {
		out.WriteString("\n")
		out.WriteString(c[i].SQL)
		if len(c)-1 > i {
			out.WriteString(",")
		}
	}
	return out.String()
}
