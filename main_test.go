package main

import "testing"

func Test_cleanup(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				input: `"C:\Users\user\Desktop\文書1.docx"`,
			},
			want: `C:\Users\user\Desktop\文書1.docx`,
		},
		{
			name: "case2",
			args: args{
				input: " \te8450dd6f908b23c9cbd6011fe3d940b24c0420a208d6924e2d920f92c894a96\t ",
			},
			want: `e8450dd6f908b23c9cbd6011fe3d940b24c0420a208d6924e2d920f92c894a96`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanup(tt.args.input); got != tt.want {
				t.Errorf("cleanup() = %v, want %v", got, tt.want)
			}
		})
	}
}
