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
		"valid one":           {input: "help", want: []string{}},
		"valid one uppercase": {input: "HELP", want: []string{}},
		//"valid many":   {input: "store hand_soap dettol 1", want: []string{"hand_soap", "dettol", "1"}},
		"invalid one":  {input: "invalidCommand", want: []string{}},
		"invalid many": {input: "invalid commands", want: []string{}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, _ := validateCommand(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("Unequal slices\n\texpected: %v\n\tgot: %v", got, c.want)
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
				t.Errorf("Unequal slices\n\texpected: %v\n\tgot: %v", got, c.want)
			}
		})
	}
}
