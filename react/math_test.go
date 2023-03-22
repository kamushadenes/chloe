package react

import (
	"context"
	"testing"

	"github.com/kamushadenes/chloe/structs"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		content  string
		expected string
		wantErr  bool
	}{
		{
			content:  "1+1",
			expected: "2",
			wantErr:  false,
		},
		{
			content:  "dasd809qd09qdaus0d",
			expected: "",
			wantErr:  true,
		},
		{
			content:  "1+1",
			expected: "2",
			wantErr:  false,
		},
		{
			content:  "2 + 3 * 4",
			expected: "14",
			wantErr:  false,
		},
		{
			content:  "((5 + 3) * 2) / 4",
			expected: "4",
			wantErr:  false,
		},
		{
			content:  "2 ** 4",
			expected: "16",
			wantErr:  false,
		},
		{
			content:  "(6 - 4) * 3 + 7 / 2",
			expected: "9.5",
			wantErr:  false,
		},
	}
	for _, tt := range tests {

		w := BytesWriter{}
		req := structs.CalculationRequest{
			Content: tt.content,
			Writer:  &w,
		}

		t.Run(tt.content, func(t *testing.T) {
			err := Calculate(context.Background(), &req)

			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil) && string(w.Bytes) != tt.expected {
				t.Errorf("Calculate() = %v, want %v", string(w.Bytes), tt.expected)
			}
		})
	}
}
