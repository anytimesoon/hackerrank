package GradingStudents

import (
	"reflect"
	"testing"
)

func Test_gradingStudents(t *testing.T) {
	type args struct {
		grades []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "First",
			args: args{
				[]int32{98, 12, 63, 61, 66},
			},
			want: []int32{100, 12, 65, 61, 66},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gradingStudents(tt.args.grades); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gradingStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
