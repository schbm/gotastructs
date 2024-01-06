package pm

import "testing"

func TestBruteForceMatch(t *testing.T) {
	type args struct {
		t string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "correct",
			args: args{
				t: "hmarcelst",
				p: "marcel",
			},
			want: 1,
		},
		{
			name: "incorrect",
			args: args{
				t: "hmarcelst",
				p: "nina",
			},
			want: -1,
		},
		{
			name: "correct prefix",
			args: args{
				t: "hmarcelst",
				p: "hma",
			},
			want: 0,
		},
		{
			name: "correct suffix",
			args: args{
				t: "hmarcelst",
				p: "rcelst",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteForceMatch(tt.args.t, tt.args.p); got != tt.want {
				t.Errorf("BruteForceMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
