package TimeConversion

import "testing"

func Test_timeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First",
			args: args{"12:01:00AM"},
			want: "00:01:00",
		}, {
			name: "Second",
			args: args{"12:01:00PM"},
			want: "12:01:00",
		}, {
			name: "Third",
			args: args{"01:01:00AM"},
			want: "01:01:00",
		}, {
			name: "second",
			args: args{"04:01:00PM"},
			want: "16:01:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
