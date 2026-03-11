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

func Test_search(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "clean file has no print debug",
			path: "testdata/clean/add.php",
			want: false,
		},
		{
			name: "dirty file with var_export",
			path: "testdata/dirty/mul.php",
			want: true,
		},
		{
			name: "dirty file with print_r, var_dump, echo",
			path: "testdata/dirty/sub.php",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := search(tt.path)
			if err != nil {
				t.Fatalf("search(%s) returned error: %v", tt.path, err)
			}
			if got != tt.want {
				t.Errorf("search(%s) = %t, want %t", tt.path, got, tt.want)
			}
		})
	}
}

func Test_walk(t *testing.T) {
	tests := []struct {
		name string
		root string
		want bool
	}{
		{
			name: "clean directory has no print debug",
			root: "testdata/clean",
			want: false,
		},
		{
			name: "dirty directory has print debug",
			root: "testdata/dirty",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := walk(tt.root)
			if err != nil {
				t.Fatalf("walk(%s) returned error: %v", tt.root, err)
			}
			if got != tt.want {
				t.Errorf("walk(%s) = %t, want %t", tt.root, got, tt.want)
			}
		})
	}
}
