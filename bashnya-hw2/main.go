package main

import (
	"fmt"
	"strings"
	"unicode"
)

const MAX_NUMBER int = 12307
const MIN_NUMBER int = -12307
const ENTER_NUMBER_MSG string = "Введите число (от -12307 до 12307): "
const INVALID_NUMBER_MSG string = "Введено неверное значение или выход за диапазон!"
const SERVICE_ERROR_MSG string = "service error"

var (
	units       = []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	unitsFemale = []string{"", "одна", "две", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	teens       = []string{"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens        = []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds    = []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
)

func readNumber() (int, error) {
	var number int
	fmt.Print(ENTER_NUMBER_MSG)
	_, err := fmt.Scan(&number)
	return number, err
}

func plural(n int, one, few, many string) string {
	if n%100 >= 11 && n%100 <= 19 {
		return many
	}
	switch n % 10 {
	case 1:
		return one
	case 2, 3, 4:
		return few
	default:
		return many
	}
}

func digitsToWords(n int, isFemaleGender bool) string {
	if n == 0 {
		return ""
	}

	var parts []string

	// Сотни
	h := n / 100
	if h > 0 {
		parts = append(parts, hundreds[h])
	}

	// Десятки и единицы
	t_u := n % 100 // tens and units
	t := t_u / 10
	u := t_u % 10

	if t_u >= 10 && t_u <= 19 {
		// 10-19
		parts = append(parts, teens[t_u-10])
	} else {
		// 20-99
		if t > 0 {
			parts = append(parts, tens[t])
		}
		if u > 0 {
			if isFemaleGender {
				parts = append(parts, unitsFemale[u])
			} else {
				parts = append(parts, units[u])
			}
		}
	}
	return strings.Join(parts, " ")
}

func writeNumberWithWords(number int) string {
	if number == 0 {
		return "Ноль"
	}

	if number < 0 {
		return "Минус " + writeNumberWithWords(number*-1)
	}

	var resultParts []string

	thousands := number / 1000
	remainder := number % 1000

	if thousands > 0 {
		// "тысяча" - женский род
		thousandWords := digitsToWords(thousands, true)
		thousandSuffix := plural(thousands, "тысяча", "тысячи", "тысяч")
		resultParts = append(resultParts, thousandWords, thousandSuffix)
	}

	if remainder > 0 {
		// "один" - мужской род (по умолчанию)
		remainderWords := digitsToWords(remainder, false)
		resultParts = append(resultParts, remainderWords)
	}

	fullStr := strings.Join(resultParts, " ")

	// Делаем первую букву заглавной
	r := []rune(fullStr)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func main() {
	var number int = 0

	number, err := readNumber()

	for number < MIN_NUMBER || number > MAX_NUMBER || err != nil {
		fmt.Println(INVALID_NUMBER_MSG)
		number, err = readNumber()
	}

	fmt.Println("Вы ввели число: ", number)

	for number < MAX_NUMBER {
		if number < 0 {
			// Если число меньше 0 → умножьте его на -1
			number *= -1
		} else if number%7 == 0 {
			// Если число кратно 7 → умножьте его на 39
			number *= 39
		} else if number%9 == 0 {
			// Если число кратно 9 → умножьте его на 13, прибавить 1 и перейдите к следующей итерации цикла
			number = number*13 + 1
			continue
		} else {
			// Во всех остальных случаях → прибавьте 2 и умножьте на 3
			number = (number + 2) * 3
		}

		if number%(13*9) == 0 { // (13*9 = 117)
			fmt.Println(SERVICE_ERROR_MSG)
			return // Выходим из main
		}

		number++
	}

	fmt.Println()
	fmt.Printf("Итоговый результат вычислений (число): %d\n", number)

	words := writeNumberWithWords(number)
	fmt.Printf("Итоговый результат вычислений (прописью): %s\n", words)

	fmt.Println()
	fmt.Println("Завершение программы.")
}
