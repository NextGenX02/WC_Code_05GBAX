package utils

import (
	"bufio"
	"ddp/database"
	_type "ddp/type"
	"fmt"
	"os"
	"strings"
)

// SearchItems Function Handle for searching items inside the database,
// if successful will return code 0 otherwise a code 1 is returned along with the error code
func SearchItems() (int, error) {
	var temp []_type.ItemsJson
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("┌───────────────────────────────────────────────┐")
	fmt.Println("│                                               │")
	fmt.Println("│   Cari Barang                                 │")
	fmt.Println("│                                               │")
	fmt.Println("└───────────────────────────────────────────────┘")
	fmt.Println("Masukan Nama barang")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	code, data, err := database.LoadDatabase()
	if err != nil {
		return 1, err
	}
	if code == 2 {
		fmt.Println("Opps...., ngak ada database nya shiku :/")
		return 0, nil
	}
	for i := 0; i < len(data); i++ {
		if strings.Contains(strings.ToLower(data[i].ItemNames), strings.ToLower(input)) {
			temp = append(temp, data[i])
		}
	}
	if len(temp) > 0 {
		fmt.Println("Barang ditemukan!:")
		for a := 0; a < len(temp); a++ {
			fmt.Printf("%d. %s | %d | %d\n", a+1, temp[a].ItemNames, temp[a].ItemPrices, temp[a].ItemStock)
		}
		return 0, nil
	}
	fmt.Printf("Barang yang bernama %s tidak ditemukan di database", input)
	return 0, nil
}
