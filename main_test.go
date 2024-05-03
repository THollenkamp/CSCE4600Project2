package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

var systemTime time.Duration

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
		{
			name: "echo command",
			args: args{
				r: strings.NewReader("echo hello\nexit\n"),
			},
			wantW: "hello",
		},
		{
			name: "help command",
			args: args{
				r: strings.NewReader("help type\nexit\n"),
			},
			wantW: "type: get file type of an input file",
		},
		{
			name: "type command",
			args: args{
				r: strings.NewReader("type main.go\nexit\n"),
			},
			wantW: "File type of main.go: go",
		},
		{
			name: "times command",
			args: args{
				r: strings.NewReader("times\nexit\n"),
			},
			wantW: "sys\t" + systemTime.String() + "\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}
