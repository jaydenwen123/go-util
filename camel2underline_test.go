package util

import "testing"

func TestCamel2Underline(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				name: "UICross",
			},
			want: "u_i_cross",
		},
		{
			name: "2",
			args: args{
				name: "DaStat",
			},
			want: "da_stat",
		},
		{
			name: "3",
			args: args{
				name: "Da",
			},
			want: "da",
		},
		{
			name: "4",
			args: args{
				name: "ExtInfo",
			},
			want: "ext_info",
		},{
			name: "5",
			args: args{
				name: "sdfUIsdfCross",
			},
			want: "sdf_u_isdf_cross",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Came2Snake(tt.args.name); got != tt.want {
				t.Errorf("Camel2Underline() = %v, want %v", got, tt.want)
			} else {
				t.Log("the got:", got, ",the want:", tt.want)
			}
		})
	}
}

