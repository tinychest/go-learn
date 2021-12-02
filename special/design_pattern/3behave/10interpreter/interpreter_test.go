package _0interpreter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlterRule_Interpreter(t *testing.T) {
	tests := []struct {
		name string
		args map[string]float64
		rule string
		want bool
	}{
		{
			name: "test01",
			rule: "a > 1 && b > 2 && c > 3",
			args: map[string]float64{
				"a": 2,
				"b": 3,
				"c": 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewAlterRule(tt.rule)
			require.NoError(t, err)
			if got := r.Interpreter(tt.args); got != tt.want {
				t.Errorf("Interpreter() = %v, want %v", got, tt.want)
			}
		})
	}
}
