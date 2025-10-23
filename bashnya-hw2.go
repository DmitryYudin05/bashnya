package main

import (
	"fmt"
)

func ChisloDoDesati(a int) string {
	b := ""
	switch a {
	case 1:
		b += "один"
	case 2:
		b += "два"
	case 3:
		b += "три"
	case 4:
		b += "четыре"
	case 5:
		b += "пять"
	case 6:
		b += "шесть"
	case 7:
		b += "семь"
	case 8:
		b += "восемь"
	case 9:
		b += "девять"
	}
	return b
}
func ChisloDoSta(a int) string {
	b := ""
	switch a / 10 {
	case 0:
		b += ChisloDoDesati(a % 10)
	case 1:
		switch a % 10 {
		case 0:
			b += "десять"
		case 1, 3:
			b += ChisloDoDesati(a%10) + "надцать"
		case 2:
			b += "двенадцать"
		case 4, 5, 6, 7, 8, 9:
			b = ChisloDoDesati(a % 10)
			b = b[0:len(b)-2] + "надцать"
		}
	case 2, 3:
		b += ChisloDoDesati(a/10) + "дцать " + ChisloDoDesati(a%10)
	case 4:
		b += "сорок " + ChisloDoDesati(a%10)
	case 5, 6, 7, 8:
		b += ChisloDoDesati(a/10) + "десят " + ChisloDoDesati(a%10)
	case 9:
		b += "девяносто " + ChisloDoDesati(a%10)
	}
	return b
}
func ChisloDoTysachi(a int) string {
	b := ""
	switch a / 100 {
	case 0:
		b += ChisloDoSta(a % 100)
	case 1:
		b += "сто " + ChisloDoSta(a%100)
	case 2:
		b += "двести " + ChisloDoSta(a%100)
	case 3, 4:
		b += ChisloDoDesati(a/100) + "ста " + ChisloDoSta(a%100)
	case 5, 6, 7, 8, 9:
		b += ChisloDoDesati(a/100) + "сот " + ChisloDoSta(a%100)
	}
	return b
}
func ChisloDoMilliona(a int) string {
	b := ""
	switch (a / 1000) % 10 {
	case 1:
		b = ChisloDoTysachi(a / 1000)
		b = b[0:len(b)-8] + "одна тысяча " + ChisloDoTysachi(a%1000)
	case 2, 3, 4:
		b += ChisloDoTysachi(a/1000) + " тысячи " + ChisloDoTysachi(a%1000)
	case 0, 5, 6, 7, 8, 9:
		b += ChisloDoTysachi(a/1000) + " тысяч " + ChisloDoTysachi(a%1000)
	}
	return b
}
func main() {
	chislo := 0
	oshib := false
	fmt.Print("Введите число: ")
	fmt.Scan(&chislo)
	if chislo < 12307 {
		for chislo < 12307 {
			switch {
			case chislo < 0:
				chislo *= -1
			case chislo%7 == 0:
				chislo *= 39
			case chislo%9 == 0:
				chislo = chislo*13 + 1
				continue
			default:
				chislo = (chislo + 2) * 3
			}
			if (chislo%13 == 0) && (chislo%9 == 0) {
				fmt.Println("\"service error\"")
				oshib = true
				break
			} else {
				chislo += 1
			}
		}
		if !oshib {
			fmt.Printf("В результате преобразований получили число %d\n", chislo)
			fmt.Printf("Результат в прописном формате: %s\n", ChisloDoMilliona(chislo))
		}
	} else {
		fmt.Printf("Число %d не меньше 12307\n", chislo)
	}
}
