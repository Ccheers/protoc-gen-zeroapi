package main

import "testing"

func Test_goPackage(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				dir: "/Users/eric/GoProject/protoc-gen-zeroapi/example/api/google",
			},
			want: "github.com/Ccheers/protoc-gen-zeroapi/example/api/google",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goPackage(tt.args.dir); got != tt.want {
				t.Errorf("goPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}
