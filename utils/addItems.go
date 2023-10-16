package utils

import (
	"bufio"
	"ddp/database"
	_type "ddp/type"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// InputItems Function that handles item input into database from the user,
// if successful will return code 0 otherwise a code 1 is returned along with the error code
func InputItems() (int, error) {
	items := _type.Items{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("┌───────────────────────────────────────────────┐")
	fmt.Println("│                                               │")
	fmt.Println("│   Input Barang                                │")
	fmt.Println("│                                               │")
	fmt.Println("└───────────────────────────────────────────────┘")
	fmt.Println("Nama Barang:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	items.ItemNames = input
	fmt.Println("Harga Barang:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	itemPrice, parsePriceError := strconv.Atoi(input)
	if parsePriceError != nil {
		return 1, parsePriceError
	}
	items.ItemPrices = itemPrice
	fmt.Println("Stok Barang:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	itemStock, parseStockError := strconv.Atoi(input)
	if parseStockError != nil {
		return 1, parseStockError
	}
	items.ItemStock = itemStock
	_, saveError := database.SaveData(&items)
	if saveError != nil {
		return 1, saveError
	}
	fmt.Printf("Barang %s Telah di simpan!", items.ItemNames)
	return 0, nil
}
