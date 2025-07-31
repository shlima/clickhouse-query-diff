package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelect_Columns(t *testing.T) {
	t.Parallel()

	type args struct {
		input string
		want  Columns
	}

	tests := []args{
		{
			input: "select 1 AS name",
			want: Columns{
				{
					Alias: "name",
					SQL:   "1 AS name",
				},
			},
		},
		{
			input: "select 1,2;",
			want: Columns{
				{
					Alias: "1",
					SQL:   "1",
				},
				{
					Alias: "2",
					SQL:   "2",
				},
			},
		},
		{
			input: "select now();",
			want: Columns{
				{
					Alias: "now()",
					SQL:   "now()",
				},
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			p, err := Parse(tt.input)
			require.NoError(t, err)
			s, err := p.Select()
			require.NoError(t, err)
			got, err := s.Columns()
			require.NoError(t, err)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
