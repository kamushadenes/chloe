package utils

import (
	"testing"
)

func TestGetFunctionName(t *testing.T) {
	tests := []struct {
		name     string
		function interface{}
		wantName string
	}{
		{
			name:     "test GetFunctionName with a basic function",
			function: func() {},
			wantName: "github.com/kamushadenes/chloe/utils.TestGetFunctionName.func1",
		},
		{
			name:     "test GetFunctionName with a named function",
			function: namedFunction,
			wantName: "github.com/kamushadenes/chloe/utils.namedFunction",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFunctionName(tt.function); got != tt.wantName {
				t.Errorf("GetFunctionName() = %v, want %v", got, tt.wantName)
			}
		})
	}
}

func namedFunction() {}
