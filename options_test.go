package main

import (
	"reflect"
	"testing"

	"github.com/docopt/docopt-go"
)

var usageTestTable = []struct {
	argv []string
	opts Options
}{
	{
		[]string{"start", "smug"},
		Options{"start", "smug", []string{}, false},
	},
	{
		[]string{"start", "smug", "-wfoo"},
		Options{"start", "smug", []string{"foo"}, false},
	},
	{
		[]string{"start", "smug:foo,bar"},
		Options{"start", "smug", []string{"foo", "bar"}, false},
	},
	{
		[]string{"start", "smug", "--attach"},
		Options{"start", "smug", []string{}, true},
	},
}

func TestParseOptions(t *testing.T) {
	parser := docopt.Parser{}
	for _, v := range usageTestTable {
		opts, err := ParseOptions(parser, v.argv)

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(v.opts, opts) {
			t.Errorf("expected struct %v, got %v", v.opts, opts)
		}
	}
}
