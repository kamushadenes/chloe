package prompts

import (
	"github.com/kamushadenes/chloe/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrompt(t *testing.T) {
	type args struct {
		prompt string
		pa     *PromptArgs
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test get prompt with valid params",
			args: args{
				prompt: "test",
				pa: &PromptArgs{
					Args: map[string]interface{}{
						"variable_1": "value1",
						"variable_2": "value2",
					},
					Mode: "key",
				},
			},
			want:    "This is just a test value1 value2",
			wantErr: false,
		},
		{
			name: "test get prompt with invalid prompt",
			args: args{
				prompt: "invalid_prompt",
				pa: &PromptArgs{
					Args: map[string]interface{}{
						"variable_1": "value1",
						"variable_2": "value2",
					},
					Mode: "invalid_mode",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "test get prompt with blank prompt",
			args: args{
				prompt: "",
				pa: &PromptArgs{
					Args: map[string]interface{}{
						"variable_1": "value1",
						"variable_2": "value2",
					},
					Mode: "invalid_mode",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "test get prompt with nil arguments",
			args: args{
				prompt: "test",
				pa:     nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "test get prompt with nil map interface",
			args: args{
				prompt: "test",
				pa: &PromptArgs{
					Args: nil,
					Mode: "value",
				},
			},
			want:    "This is just a test <no value> <no value>",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPrompt(tt.args.prompt, tt.args.pa)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPromptSize(t *testing.T) {
	type args struct {
		prompt string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test get prompt size with valid prompt",
			args: args{
				prompt: "test",
			},
			want:    13,
			wantErr: false,
		},
		{
			name: "test get prompt size with invalid prompt",
			args: args{
				prompt: "invalid_prompt",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test get prompt size with blank prompt",
			args: args{
				prompt: "",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPromptSize(tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPromptSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPromptSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListPrompts(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		{
			name: "test list prompts",
			want: []string{"default", "bootstrap", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListPrompts()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPrompts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for k := range tt.want {
				assert.True(t, utils.StringInSlice(tt.want[k], got), "ListPrompts() = %v, want %v", got, tt.want)
			}
		})
	}
}
