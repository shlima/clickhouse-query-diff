package select_diff

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type Service struct {
	LeftSQL  string
	RightSQL string
}

func New() *Service {
	return &Service{}
}

func (s *Service) SetLeftSQL(sql string) {
	s.LeftSQL = sql
}

func (s *Service) SetRightSQL(sql string) {
	s.RightSQL = sql
}

func (s *Service) ColumnsDiffHTML() (string, error) {
	left, err := reorderSelectColumns(s.LeftSQL)
	if err != nil {
		return "", fmt.Errorf("failed to reorder left columns: %w", err)
	}

	right, err := reorderSelectColumns(s.RightSQL)
	if err != nil {
		return "", fmt.Errorf("failed to reorder right columns: %w", err)
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(left, right, false)
	return dmp.DiffPrettyHtml(diffs), nil
}
