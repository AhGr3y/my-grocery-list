package main

import (
	"testing"
)

func TestCommandValidation(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"empty":               {input: "", want: ""},
		"valid one lowercase": {input: "help", want: "help"},
		"valid one uppercase": {input: "HELP", want: "help"},
		"many":                {input: "store hand_soap dettol 1", want: ""},
		"invalid one":         {input: "invalidCommand", want: ""},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, _ := validateCommand(c.input)
			if got != c.want {
				t.Errorf("Unequal slices\n\texpected: %v\n\tgot: %v", c.want, got)
			}
		})
	}
}

func TestProcessYesNo(t *testing.T) {
	cases := map[string]struct {
		input       string
		want_result bool
		want_err    error
	}{
		"empty":     {input: "", want_result: true, want_err: nil},
		"y lower":   {input: "y", want_result: true, want_err: nil},
		"y upper":   {input: "Y", want_result: true, want_err: nil},
		"yes lower": {input: "yes", want_result: true, want_err: nil},
		"yes upper": {input: "YES", want_result: true, want_err: nil},
		"n lower":   {input: "n", want_result: false, want_err: nil},
		"n upper":   {input: "N", want_result: false, want_err: nil},
		"no lower":  {input: "no", want_result: false, want_err: nil},
		"no upper":  {input: "NO", want_result: false, want_err: nil},
		"invalid":   {input: "maybe", want_result: false, want_err: ErrInvalidCommand},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got_result, got_err := processYesNo(c.input)
			if got_result != c.want_result {
				t.Errorf("Unexpected result\n\texpected: %v\n\tgot: %v", c.want_result, got_result)
			}
			if got_err != c.want_err {
				t.Errorf("Unexpected result\n\texpected: %v\n\tgot: %v", c.want_err, got_err)
			}
		})
	}
}
