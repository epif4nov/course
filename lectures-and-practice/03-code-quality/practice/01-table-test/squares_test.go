package tabletest

import (
	"errors"
	"testing"
)

func TestSumSquares(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
		wantErr error
	}{
		{
			name:    "ordinary input",
			numbers: []int{1, 2, 3},
			want:    14,
		},
		{
			name:    "empty input",
			numbers: []int{},
			want:    0,
		},
		{
			name:    "negative number",
			numbers: []int{2, -3, 4},
			wantErr: ErrNegativeNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumSquares(tt.numbers)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("SumSquares() error = %v, want error wrapping %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("SumSquares() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Errorf("SumSquares() = %d, want %d", got, tt.want)
			}
		})
	}
}
