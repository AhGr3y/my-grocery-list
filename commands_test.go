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
