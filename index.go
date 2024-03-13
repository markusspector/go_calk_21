//попытка номер 21

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// мапа с римскими числами.
var mapRomanian = map[string]int{
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
}

// функция проверки существования римских чисел

func RomanianToArabicFunc(x string) bool {
	_, exists := mapRomanian[x]
	return exists
}

// получение арабского числа из мапы по римскому

func ArabicToRomanian(m map[string]int, value int) string {

	for key, val := range m {
		if val == value {
			return key
		}
	}

	return ""
}

// получение римского числа из мапы по арабскому

func RomanianToArabic(x string) int {
	return mapRomanian[x]
}

func main() {

	var arg1 int
	var arg2 string
	var arg3 int
	var result int
	rom := false // вводятся римские числа?

	fmt.Println("Калькулятор:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	// Удаляем символ новой строки из конца введенной строки
	input = strings.TrimSpace(input)

	// Разбиваем строку на элементы по пробелу
	elements := strings.Split(input, " ")

	if len(elements) > 3 {
		panic("error количество операндов должно быть 2")
	}

	for i, elem := range elements {

		if i == 0 {
			if RomanianToArabicFunc(elem) {
				arg1 = RomanianToArabic(elem)
				if arg1 > 10 {
					panic("error числа должны быть не более X")
				}
				rom = true
			} else {
				arg1, err = strconv.Atoi(elem)
				if err != nil || arg1 > 10 {
					panic("error")
				}
			}
		}

		if i == 1 {
			arg2 = elem
		}

		if i == 2 {
			if RomanianToArabicFunc(elem) {
				if !rom {
					panic("error оба числа должны быть римские или арабские одновременно")
				}
				if arg1 > 10 {
					panic("error числа должны быть не более X")
				}
				arg3 = RomanianToArabic(elem)
			} else {
				arg3, err = strconv.Atoi(elem)

				if err != nil || arg3 > 10 || rom {
					panic("error")
				}
			}
		}
	}

	switch arg2 {
	case "+":
		result = arg1 + arg3
	case "-":
		result = arg1 - arg3
	case "*":
		result = arg1 * arg3
	case "/":
		{
			if arg3 == 0 {
				panic("На ноль делить нельзя")
			} else {
				result = arg1 / arg3
			}

		}
	default:
		panic("error")
	}

	if rom {
		if result < 1 {
			panic("error")
		}

		fmt.Println(ArabicToRomanian(mapRomanian, result))
	} else {
		fmt.Println(result)
	}
}
