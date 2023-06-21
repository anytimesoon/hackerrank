package PlusMinus

import (
	"io"
	"os"
	"testing"
)

func TestPlusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First",
			args: args{
				arr: []int32{-4, 3, -9, 0, 4, 1},
			},
			want: "0.500000\n0.333333\n0.166667\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PlusMinus(tt.args.arr)

			err := w.Close()
			if err != nil {
				return
			}

			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			if string(out) != tt.want {
				t.Errorf("Expected %s, got %s", tt.want, out)
			}
		})
	}
}
