package FormingMagicSquare

import "testing"

func TestFormingMagicSquare(t *testing.T) {
	type args struct {
		s [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "First",
			args: args{s: [][]int32{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}}},
			want: 4,
		}, {
			name: "Second",
			args: args{s: [][]int32{{8, 3, 6}, {1, 5, 9}, {6, 2, 2}}},
			want: 7,
		},
		{
			name: "Third",
			args: args{s: [][]int32{{6, 7, 2}, {4, 5, 9}, {8, 3, 4}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormingMagicSquare(tt.args.s); got != tt.want {
				t.Errorf("FormingMagicSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
