package html

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func TestCodeToHtml(t *testing.T) {
	type args struct {
		in0 *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{in0: &cli.Context{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CodeToHtml(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("CodeToHtml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
