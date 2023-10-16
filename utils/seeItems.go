package utils

import (
	"ddp/database"
	"fmt"
)

// SeeItems Function that handles a list of items that currently save
// inside a database,if successful will return code 0 otherwise a code 1 is returned along with the error code
func SeeItems() (int, error) {
	fmt.Println("┌───────────────────────────────────────────────┐")
	fmt.Println("│                                               │")
	fmt.Println("│   Daftar Barang                               │")
	fmt.Println("│                                               │")
	code, data, err := database.LoadDatabase()
	if err != nil {
		return 1, err
	}
	if code == 2 {
		fmt.Println("│ Tidak ada list barang atau Shiku Database tidak ditemukan!")
		fmt.Println("└───────────────────────────────────────────────┘")
		return 0, nil
	}
	for i := 0; i < len(data); i++ {
		fmt.Printf("│   %d. %s | %d | %d   │\n", i+1, data[i].ItemNames, data[i].ItemPrices, data[i].ItemStock)
	}
	fmt.Println("└───────────────────────────────────────────────┘")
	return 0, nil
}
