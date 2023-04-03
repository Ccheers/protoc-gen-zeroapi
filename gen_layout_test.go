package main

import "testing"

func Test_genLayout(t *testing.T) {
	type args struct {
		outDir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				outDir: "/Users/eric/GoProject/protoc-gen-zeroapi/internal/api",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genZeroLayout(tt.args.outDir)
		})
	}
}
