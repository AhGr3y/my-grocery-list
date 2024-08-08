package main

import (
	"reflect"
	"testing"
)

func TestCommandValidation(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []string
	}{
		"empty":               {input: "", want: []string{}},
		"valid one":           {input: "help", want: []string{"help"}},
		"valid one uppercase": {input: "HELP", want: []string{"help"}},
		"valid many":          {input: "store hand_soap dettol 1", want: []string{"store", "hand_soap", "dettol", "1"}},
		"invalid one":         {input: "invalidCommand", want: []string{}},
		"invalid many":        {input: "invalid commands", want: []string{}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, _ := validateCommand(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("Unequal slices\n\texpected: %v\n\tgot: %v", c.want, got)
			}
		})
	}
}

func TestWordsToLower(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  []string
	}{
		"empty":            {input: []string{}, want: []string{}},
		"one word lower":   {input: []string{"hello"}, want: []string{"hello"}},
		"one word upper":   {input: []string{"HELLO"}, want: []string{"hello"}},
		"many words lower": {input: []string{"hello", "world"}, want: []string{"hello", "world"}},
		"many words upper": {input: []string{"HELLO", "WORLD"}, want: []string{"hello", "world"}},
		"many characters":  {input: []string{"Hello", "1", "$"}, want: []string{"hello", "1", "$"}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := wordsToLower(c.input)
			if !reflect.DeepEqual(got, c.want) {
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
