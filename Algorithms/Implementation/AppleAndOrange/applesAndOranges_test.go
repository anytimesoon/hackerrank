package AppleAndOrange

import (
	"reflect"
	"testing"
)

func Test_checkIfInRange(t *testing.T) {
	type args struct {
		s    int32
		t    int32
		dist int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{4, 8, 6},
			want: true,
		}, {
			name: "Second",
			args: args{4, 8, 8},
			want: true,
		}, {
			name: "Third",
			args: args{4, 8, 4},
			want: true,
		}, {
			name: "Fifth",
			args: args{4, 8, 2},
			want: false,
		}, {
			name: "Sixth",
			args: args{4, 8, 10},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfInRange(tt.args.s, tt.args.t, tt.args.dist); got != tt.want {
				t.Errorf("checkIfInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountApplesAndOranges(t *testing.T) {
	type args struct {
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "First",
			args: args{
				s:       7,
				t:       11,
				a:       5,
				b:       15,
				apples:  []int32{-2, 2, 1},
				oranges: []int32{5, -6},
			},
			want: []int32{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountApplesAndOranges(tt.args.s, tt.args.t, tt.args.a, tt.args.b, tt.args.apples, tt.args.oranges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountApplesAndOranges() = %v, want %v", got, tt.want)
			}
		})
	}
}
