package select_diff

import (
	"fmt"
	"sort"

	"github.com/shlima/clickhouse-query-diff/internal/pkg/parser"
)

func reorderSelectColumns(sql string) (string, error) {
	p, err := parser.Parse(sql)
	if err != nil {
		return "", fmt.Errorf("failed to parse: %w", err)
	}

	s, err := p.Select()
	if err != nil {
		return "", fmt.Errorf("failed to extract select: %w", err)
	}

	columns, err := s.Columns()
	if err != nil {
		return "", fmt.Errorf("failed to extract columns: %w", err)
	}

	sort.Slice(columns, func(i, j int) bool {
		return columns[i].Alias < columns[j].Alias
	})

	return "SELECT " + columns.String(), nil
}
