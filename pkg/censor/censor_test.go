package censor

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid qwerty",
			args: args{
				content: "some qwerty text",
			},
			want: false,
		},
		{
			name: "invalid йцукен",
			args: args{
				content: "you are йцукен",
			},
			want: false,
		},
		{
			name: "invalid zxvbnm",
			args: args{
				content: "zxvbnm not dead",
			},
			want: false,
		},
		{
			name: "valid",
			args: args{
				content: "qwert йцуке xvbnm",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.content); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
