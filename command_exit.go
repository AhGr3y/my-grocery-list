package main

import "os"

func commandExit(cfg *config, args ...string) error {
	os.Exit(1)
	return nil
}
