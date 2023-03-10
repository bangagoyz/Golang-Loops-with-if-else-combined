package main

import "fmt"

func main() {
	var sym = [7]string{"\u0421", "\u0410", "\u0428", "\u0410", "\u0420", "\u0412", "\u041E"}
	var k int = 0
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
		if i == 4 {
			for j := 0; j < 11; j++ {
				fmt.Println("Nilai j = ", j)
				if j == 4 {
					for j := 0; j < len(sym); j++ {
						fmt.Println("character ", sym[j], " starts at byte position ", k)
						k = k + 2
					}
				}
			}
		}

	}

}
