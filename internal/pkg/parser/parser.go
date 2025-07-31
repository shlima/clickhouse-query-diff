package parser

import (
	"fmt"

	clickhouse "github.com/AfterShip/clickhouse-sql-parser/parser"
)

type Parser struct {
	statement clickhouse.Expr
}

func allocate(statement clickhouse.Expr) *Parser {
	return &Parser{statement: statement}
}

func Parse(sql string) (*Parser, error) {
	parser := clickhouse.NewParser(sql)
	statements, err := parser.ParseStmts()
	switch {
	case err != nil:
		return nil, fmt.Errorf("failed to parse statements: %w", err)
	case len(statements) == 0:
		return nil, ErrNoSelect
	case len(statements) > 1:
		return nil, fmt.Errorf("%s %d", ErrMultipleSelects, len(statements))
	default:
		return allocate(statements[0]), nil
	}
}

func (p *Parser) Select() (*Select, error) {
	switch v := p.statement.(type) {
	case *clickhouse.SelectQuery:
		return newSelect(v), nil
	case *clickhouse.CreateMaterializedView:
		return newSelect(v.SubQuery.Select), nil
	default:
		return nil, fmt.Errorf("i don't know about type %T", v)
	}
}
