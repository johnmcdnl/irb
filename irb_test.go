package irb

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		home     float64
		away     float64
		handicap float64
		sa       Result
		sb       Result
	}
	tests := []struct {
		name    string
		args    args
		want    *IRB
		wantErr bool
	}{
		{"test", args{home: 90.09, away: 85.02, sa: Win, sb: Loose}, &IRB{}, false},
		{"test", args{home: 90.09, away: 85.02, sa: Loose, sb: Win}, &IRB{}, false},
		{"test", args{home: 90.09, away: 85.02, sa: Draw, sb: Draw}, &IRB{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.home, tt.args.away, tt.args.sa, tt.args.sb)
		})
	}
}
