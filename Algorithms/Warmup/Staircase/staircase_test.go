package Staircase

import (
	"io"
	"os"
	"testing"
)

func TestStaircase(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First",
			args: args{
				n: 6,
			},
			want: "     #\n    ##\n   ###\n  ####\n #####\n######\n",
		}, {
			name: "Second",
			args: args{
				n: 1,
			},
			want: "#\n",
		}, {
			name: "Third",
			args: args{
				n: 0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			Staircase(tt.args.n)

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
