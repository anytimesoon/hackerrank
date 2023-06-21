package Mini_MaxSum

import "testing"

func TestMiniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "First",
			args: args{[]int32{23, 8745, 383, 93, 9}},
			want: []int64{508, 9244},
		}, {
			name: "Second",
			args: args{[]int32{45, 2344, 23, 4, 567}},
			want: []int64{639, 2979},
		}, {
			name: "Third",
			args: args{[]int32{897689, 34534343, 345344534, 434345, 567}},
			want: []int64{35866944, 381210911},
		}, {
			name: "Forth",
			args: args{[]int32{256741038, 623958417, 467905213, 714532089, 938071625}},
			want: []int64{2063136757, 2744467344},
		}, {
			name: "Fifth",
			args: args{[]int32{942381765, 627450398, 954173620, 583762094, 236817490}},
			want: []int64{2390411747, 3107767877},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MiniMaxSum(tt.args.arr)
			if output[0] != tt.want[0] || output[1] != tt.want[1] {
				t.Errorf("expected %v, got %v", tt.want, output)
			}
		})
	}
}
