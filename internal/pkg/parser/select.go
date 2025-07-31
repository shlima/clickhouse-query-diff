package parser

import (
	"fmt"

	clickhouse "github.com/AfterShip/clickhouse-sql-parser/parser"
)

type Select struct {
	statement *clickhouse.SelectQuery
}

func newSelect(statement *clickhouse.SelectQuery) *Select {
	return &Select{statement: statement}
}

func (s *Select) Columns() (Columns, error) {
	columns := make([]Column, 0)
	for _, expr := range s.statement.SelectColumns.Items {
		alias, err := columnAlias(expr)
		if err != nil {
			return nil, fmt.Errorf("failed to extract column alias: %w", err)
		}

		columns = append(columns, Column{
			Alias: alias,
			SQL:   expr.String(0),
		})
	}

	return columns, nil
}

func columnAlias(input clickhouse.Expr) (string, error) {
	switch v := input.(type) {
	case *clickhouse.AliasExpr:
		return v.Alias.String(0), nil
	case *clickhouse.NumberLiteral:
		return v.String(0), nil
	case *clickhouse.FunctionExpr:
		return v.String(0), nil
	default:
		return "", fmt.Errorf("i don't know column alias %T", input)
	}
}
