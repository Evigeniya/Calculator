package main

import (
	"fmt"
)

func main() {
	var a, b, c, num string
	var x, y, sum int
	for {
		rim := 0
		p := 0
		fmt.Scanln(&a, &num, &b, &c)
		if c != "" {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			break
		}

		if a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9" || a == "10" {
			p = 1
		}
		if b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" || b == "10" {
			p += 1
		}
		if a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX" || a == "X" {
			rim = 1
		}
		if b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX" || b == "X" {
			rim += 1
		}

		if p == 2 || rim == 2 {
			switch {
			case a == "1" || a == "I":
				x = 1
			case a == "2" || a == "II":
				x = 2
			case a == "3" || a == "III":
				x = 3
			case a == "4" || a == "IV":
				x = 4
			case a == "5" || a == "V":
				x = 5
			case a == "6" || a == "VI":
				x = 6
			case a == "7" || a == "VII":
				x = 7
			case a == "8" || a == "VII":
				x = 8
			case a == "9" || a == "IX":
				x = 9
			case a == "10" || a == "X":
				x = 10

			}
			switch {
			case b == "1" || b == "I":
				y = 1
			case b == "2" || b == "II":
				y = 2
			case b == "3" || b == "III":
				y = 3
			case b == "4" || b == "IV":
				y = 4
			case b == "5" || b == "V":
				y = 5
			case b == "6" || b == "VI":
				y = 6
			case b == "7" || b == "VII":
				y = 7
			case b == "8" || b == "VII":
				y = 8
			case b == "9" || b == "IX":
				y = 9
			case b == "10" || b == "X":
				y = 10

			}
			switch num {
			case "+":
				sum = x + y
			case "-":
				sum = x - y
			case "*":
				sum = x * y
			case "/":
				sum = x / y

			}
		} else if p == 1 && rim == 1 {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			break
		} else {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			break
		}

		if num != "+" && num != "-" && num != "/" && num != "*" {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			break
		} else if p == 2 {
			fmt.Println(sum)
		} else if sum < 1 {

			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			break
		} else {
			switch sum {
			case 1:
				fmt.Println("I")
			case 2:
				fmt.Println("II")
			case 3:
				fmt.Println("III")
			case 4:
				fmt.Println("IV")
			case 5:
				fmt.Println("V")
			case 6:
				fmt.Println("VI")
			case 7:
				fmt.Println("VII")
			case 8:
				fmt.Println("VIII")
			case 9:
				fmt.Println("IX")
			case 10:
				fmt.Println("X")
			case 11:
				fmt.Println("XI")
			case 12:
				fmt.Println("XII")
			case 13:
				fmt.Println("XIII")
			case 14:
				fmt.Println("XIV")
			case 15:
				fmt.Println("XV")
			case 16:
				fmt.Println("XVI")
			case 17:
				fmt.Println("XVII")
			case 18:
				fmt.Println("XVIII")
			case 19:
				fmt.Println("XIX")
			case 20:
				fmt.Println("XX")
			case 21:
				fmt.Println("XXI")
			case 22:
				fmt.Println("XXII")
			case 23:
				fmt.Println("XXIII")
			case 24:
				fmt.Println("XXIV")
			case 25:
				fmt.Println("XXV")
			case 26:
				fmt.Println("XXVI")
			case 27:
				fmt.Println("XXVII")
			case 28:
				fmt.Println("XXVIII")
			case 29:
				fmt.Println("XXIX")
			case 30:
				fmt.Println("XXX")
			case 31:
				fmt.Println("XXXI")
			case 32:
				fmt.Println("XXXII")
			case 33:
				fmt.Println("XXXIII")
			case 34:
				fmt.Println("XXXIV")
			case 35:
				fmt.Println("XXXV")
			case 36:
				fmt.Println("XXXVI")
			case 37:
				fmt.Println("XXXVII")
			case 38:
				fmt.Println("XXXVIII")
			case 39:
				fmt.Println("XXXIX")
			case 40:
				fmt.Println("XL")
			case 41:
				fmt.Println("XLI")
			case 42:
				fmt.Println("XLII")
			case 43:
				fmt.Println("XLIII")
			case 44:
				fmt.Println("XLIV")
			case 45:
				fmt.Println("XLV")
			case 46:
				fmt.Println("XLVI")
			case 47:
				fmt.Println("XLVII")
			case 48:
				fmt.Println("XLVIII")
			case 49:
				fmt.Println("XLIX")
			case 50:
				fmt.Println("L")
			case 51:
				fmt.Println("LI")
			case 52:
				fmt.Println("LII")
			case 53:
				fmt.Println("LIII")
			case 54:
				fmt.Println("LIV")
			case 55:
				fmt.Println("LV")
			case 56:
				fmt.Println("LVI")
			case 57:
				fmt.Println("LVII")
			case 58:
				fmt.Println("LVIII")
			case 59:
				fmt.Println("LIX")
			case 60:
				fmt.Println("LX")
			case 61:
				fmt.Println("LXI")
			case 62:
				fmt.Println("LXII")
			case 63:
				fmt.Println("LXIII")
			case 64:
				fmt.Println("LXIV")
			case 65:
				fmt.Println("LXV")
			case 66:
				fmt.Println("LXVI")
			case 67:
				fmt.Println("LXVII")
			case 68:
				fmt.Println("LXVIII")
			case 69:
				fmt.Println("LXIX")
			case 70:
				fmt.Println("LXX")
			case 71:
				fmt.Println("LXXI")
			case 72:
				fmt.Println("LXXII")
			case 73:
				fmt.Println("LXXIII")
			case 74:
				fmt.Println("LXXIV")
			case 75:
				fmt.Println("LXXV")
			case 76:
				fmt.Println("LXXVI")
			case 77:
				fmt.Println("LXXVII")
			case 78:
				fmt.Println("LXXVIII")
			case 79:
				fmt.Println("LXXIX")
			case 80:
				fmt.Println("LXXX")
			case 81:
				fmt.Println("LXXXI")
			case 82:
				fmt.Println("LXXXII")
			case 83:
				fmt.Println("LXXXIII")
			case 84:
				fmt.Println("LXXXIV")
			case 85:
				fmt.Println("LXXXV")
			case 86:
				fmt.Println("LXXXVI")
			case 87:
				fmt.Println("LXXXVII")
			case 88:
				fmt.Println("LXXXVIII")
			case 89:
				fmt.Println("LXXXIX")
			case 90:
				fmt.Println("XC")
			case 91:
				fmt.Println("XCI")
			case 92:
				fmt.Println("XC")
			case 93:
				fmt.Println("XCIII")
			case 94:
				fmt.Println("XCIV")
			case 95:
				fmt.Println("XCXV")
			case 96:
				fmt.Println("XCVI")
			case 97:
				fmt.Println("XCVII")
			case 98:
				fmt.Println("XCVIII")
			case 99:
				fmt.Println("XCIX")
			case 100:
				fmt.Println("C")
			}
		}
	}
}
