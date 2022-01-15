package main

import (
	"testing"
)

func Test_isPrintDebug(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "print('hoge');",
			want:  true,
		},
		{
			input: "print_r('hoge');",
			want:  true,
		},
		{
			input: "printOriginal('hoge');",
			want:  false,
		},
		{
			input: "var_dump('piyo');",
			want:  true,
		},
		{
			input: "var_export('fuga');",
			want:  true,
		},
		{
			input: "echo 'a';",
			want:  true,
		},
		{
			input: "echo('a');",
			want:  true,
		},
		{
			input: "echoOriginal('a');",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isPrintDebug(tt.input); got != tt.want {
				t.Errorf("input: %s, got: %t, want: %t", tt.input, got, tt.want)
			}
		})
	}
}
