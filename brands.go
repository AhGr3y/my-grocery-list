package main

import (
	"context"
	"strings"

	"github.com/AhGr3y/my-grocery-list/internal/database"
)

func checkBrand(cfg *config, brand string) (database.Brand, bool) {
	// brand is not case-sensitive
	brandLower := strings.ToLower(brand)

	// Retrieve brand data from database via brand name lookup
	dbBrand, err := cfg.DB.GetBrand(context.Background(), brandLower)
	if err != nil {
		return database.Brand{}, false
	}
	return dbBrand, true
}
