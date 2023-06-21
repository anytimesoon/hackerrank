package DiagonalDifference

import "testing"

func Test_diagonalDifference(t *testing.T) {
	type args struct {
		arr [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "First",
			args: args{
				arr: [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiagonalDifference(tt.args.arr); got != tt.want {
				t.Errorf("DiagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
