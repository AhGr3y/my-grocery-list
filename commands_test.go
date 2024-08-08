package main

import "testing"

func TestCommandValidation(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []string
	}{
		"empty":               {input: "", want: []string{}},
		"valid one":           {input: "help", want: []string{}},
		"valid one uppercase": {input: "HELP", want: []string{}},
		//"valid many":   {input: "store hand_soap dettol 1", want: []string{"hand_soap", "dettol", "1"}},
		"invalid one":  {input: "invalidCommand", want: []string{}},
		"invalid many": {input: "invalid commands", want: []string{}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, _, _ := validateCommand(c.input)
			if len(got) != len(c.want) {
				t.Errorf("Unexpected slice length\n\texpected: %v\n\tgot: %v", len(got), len(c.want))
			}
			for i := range got {
				if got[i] != c.want[i] {
					t.Errorf("Unexpected slice item\n\texpected: %v\n\tgot: %v", got[i], c.want[i])
				}
			}
		})
	}
}
