package main

import "testing"

func Test_helloWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TC1_OK",
			want: "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helloWorld(); got != tt.want {
				t.Errorf("helloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}
