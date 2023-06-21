package BirthdayCakeCandles

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "First",
			args: args{candles: []int32{3, 6, 2, 6, 1}},
			want: 2,
		}, {
			name: "Second",
			args: args{candles: []int32{3, 2, 1, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BirthdayCakeCandles(tt.args.candles); got != tt.want {
				t.Errorf("BirthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
