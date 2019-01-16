package closer

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type closerSuccess struct {
}

func (c closerSuccess) Close() (err error) {
	return
}

type closerError struct {
}

func (c closerError) Close() (err error) {
	err = errors.New("closer error")
	return
}

func TestCloser(t *testing.T) {

	getStdout := func(obj io.Closer) (out []byte, err error) {
		rescueStdout := os.Stdout
		defer func() { os.Stdout = rescueStdout }()
		r, w, err := os.Pipe()
		if err != nil {
			return nil, err
		}
		os.Stdout = w

		Closer(obj)

		err = w.Close()
		if err != nil {
			return
		}
		out, err = ioutil.ReadAll(r)
		return
	}

	cs := closerSuccess{}
	ce := closerError{}

	type args struct {
		body io.Closer
	}
	type expected struct {
		err bool
	}
	tests := []struct {
		name string
		args args
		want expected
	}{
		{
			name: "success",
			args: args{
				body: cs,
			},
			want: expected{
				err: false,
			},
		},
		{
			name: "error",
			args: args{
				body: ce,
			},
			want: expected{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := getStdout(tt.args.body)
			if err != nil {
				t.Error(err)
				return
			}

			if (len(out) > 0) != tt.want.err {
				fmt.Printf("out: %q\n", string(out))
				t.Errorf("closer() unexpected log %q", string(out))
			}
		})
	}
}
