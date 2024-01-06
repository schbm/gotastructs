package pm

import (
	"reflect"
	"testing"
)

func TestBoyerMooreLastOccurence(t *testing.T) {
	type args struct {
		p string
	}
	wantResult := map[string]int{
		"a": 4,
		"b": 5,
	}

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "correct",
			args: args{
				p: "ababab",
			},
			want: wantResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoyerMooreLastOccurence(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoyerMooreLastOccurence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoyerMooreMatch(t *testing.T) {
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
				t: "gabacaabadcabacabaabb",
				p: "abacab",
			},
			want: 11,
		},
		{
			name: "correct",
			args: args{
				t: "abaababacbacababcabab",
				p: "cabab",
			},
			want: 11,
		},
		{
			name: "correct",
			args: args{
				t: "abacab",
				p: "abacab",
			},
			want: 0,
		},
		{
			name: "correct",
			args: args{
				t: "abacad",
				p: "abacab",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoyerMooreMatch(tt.args.t, tt.args.p); got != tt.want {
				t.Errorf("BoyerMooreMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
