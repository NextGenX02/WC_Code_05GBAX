package main

import (
	"ddp/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	// Some Bored Stuff
	// run this bored function as separate thread
	go utils.DoMusic()

	// Check if the user run from command line
	if len(os.Args) > 0 {
		switch os.Args[1] {
		case strings.ToLower("input"):
			utils.InputItems()
			os.Exit(0)
		case strings.ToLower("list"):
			utils.SeeItems()
			os.Exit(0)
		case strings.ToLower("search"):
			utils.SearchItems()
			os.Exit(0)
		case strings.ToLower("about"):
			utils.ShowWhoMake()
			os.Exit(0)
		default:
			fmt.Println("Valid arguments are input,list,search or about")
			os.Exit(1)
		}
	}
}

func main() {
	// Run as usual
	var menuChoose string
	var afterMenuChoose string
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("│                                                │")
	fmt.Println("│   Repository Barang PT Maju Kena Mundur Kena   │")
	fmt.Println("│                                                │")
	fmt.Println("│                                                │")
	fmt.Println("│   Masukan pilihan Anda:                        │")
	fmt.Println("│                                                │")
	fmt.Println("│   1. Input Barang                              │")
	fmt.Println("│   2. Lihat Daftar Barang                       │")
	fmt.Println("│   3. Cari Barang                               │")
	fmt.Println("│   4. Tentang Aplikasi                          │")
	fmt.Println("│                                                │")
	fmt.Println("└────────────────────────────────────────────────┘")

	_, scanError := fmt.Scanln(&menuChoose)
	if scanError != nil {
		panic(scanError.Error())
	}
	cn, parseToIntError := strconv.ParseInt(menuChoose, 10, 32)
	if parseToIntError != nil {
		panic(fmt.Errorf("%s", "Your Choice is not a Number!"))
	}
	switch cn {
	case 1:
		utils.ClearScreen()
		code, Inputerr := utils.InputItems()
		if Inputerr != nil {
			panic(Inputerr.Error())
		}
		if code == 0 {
			fmt.Println("\nBalik ke menu? (yes/no):")
			fmt.Scanln(&afterMenuChoose)
			if strings.ToLower(afterMenuChoose) == "yes" {
				utils.ClearScreen()
				main()
			} else if strings.ToLower(afterMenuChoose) == "no" {
				fmt.Println("OK bye!")
				os.Exit(0)
			} else if strings.ToLower(afterMenuChoose) != "yes" || strings.ToLower(afterMenuChoose) != "no" {
				fmt.Println("OK Bye!, just use yes or no next time")
				os.Exit(0)
			}
		}
		break
	case 2:
		utils.ClearScreen()
		code, ListError := utils.SeeItems()
		if ListError != nil {
			panic(ListError.Error())
		}
		if code == 0 {
			fmt.Println("\nBalik ke menu? (yes/no):")
			fmt.Scanln(&afterMenuChoose)
			if strings.ToLower(afterMenuChoose) == "yes" {
				utils.ClearScreen()
				main()
			} else if strings.ToLower(afterMenuChoose) == "no" {
				fmt.Println("OK bye!")
				os.Exit(0)
			} else if strings.ToLower(afterMenuChoose) != "yes" || strings.ToLower(afterMenuChoose) != "no" {
				fmt.Println("OK Bye!, just use yes or no next time")
				os.Exit(0)
			}
		}
		break
	case 3:
		utils.ClearScreen()
		code, SearchError := utils.SearchItems()
		if SearchError != nil {
			panic(SearchError.Error())
		}
		if code == 0 {
			fmt.Println("\nBalik ke menu? (yes/no):")
			fmt.Scanln(&afterMenuChoose)
			if strings.ToLower(afterMenuChoose) == "yes" {
				utils.ClearScreen()
				main()
			} else if strings.ToLower(afterMenuChoose) == "no" {
				fmt.Println("OK bye!")
				os.Exit(0)
			} else if strings.ToLower(afterMenuChoose) != "yes" || strings.ToLower(afterMenuChoose) != "no" {
				fmt.Println("OK Bye!, just use yes or no next time")
				os.Exit(0)
			}
		}
		break
	case 4:
		code := utils.ShowWhoMake()
		if code == 0 {
			fmt.Println("\nBalik ke menu? (yes/no):")
			fmt.Scanln(&afterMenuChoose)
			if strings.ToLower(afterMenuChoose) == "yes" {
				utils.ClearScreen()
				main()
			} else if strings.ToLower(afterMenuChoose) == "no" {
				fmt.Println("OK bye!")
				os.Exit(0)
			} else if strings.ToLower(afterMenuChoose) != "yes" || strings.ToLower(afterMenuChoose) != "no" {
				fmt.Println("OK Bye!, just use yes or no next time")
				os.Exit(0)
			}
		}
		break
	}
}
