package vowels

import "testing"

func TestIsAlt(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "amazon",
			args: args{input: "amazon"},
			want: true,
		},
		{
			name: "apple",
			args: args{input: "apple"},
			want: false,
		},
		{
			name: "banana",
			args: args{input: "banana"},
			want: true,
		},
		{
			name: "cheeze",
			args: args{input: "cheeze"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlt(tt.args.input); got != tt.want {
				t.Errorf("IsAlt() = %v, want %v", got, tt.want)
			}
		})
	}
}
