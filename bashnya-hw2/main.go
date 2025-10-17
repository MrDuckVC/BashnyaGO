package main

import "fmt"

const MAX_NUMBER int = 12307
const MIN_NUMBER int = -12307
const ENTER_NUMBER_MSG string = "Введите число: "
const INVALID_NUMBER_MSG string = "Введено неверное значение или выход за диапазон!"
const SERVICE_ERROR_MSG string = "service error"

func readNumber() (int, error) {
	var number int
	fmt.Print(ENTER_NUMBER_MSG)
	_, err := fmt.Scan(&number)
	return number, err
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
		fmt.Println("Вы ввели число: ", number)
		if number < 0 {
			// Если число меньше 0 → умножьте его на -1
			number *= -1
		} else if number%7 == 0 {
			// Если число кратно 7 → умножьте его на 39
			number *= 39
		} else if number%9 == 0 {
			// Если число кратно 9 → умножьте его на 13, прибавить 1 и перейдите к следующей итерации цикла
			number *= 13
			number++
			continue
		} else {
			// Во всех остальных случаях → прибавьте 2 и умножьте на 3
			number += 2
			number *= 3
		}

		// Если число одновременно кратно 13 и 9 → завершите цикл и выведите в stdout(консоль/терминал) строку: "service error"
		if number%(13*9) == 0 {
			fmt.Println(SERVICE_ERROR_MSG)
			break
		}
		// Иначе → прибавьте к числу 1
		number++
	}

	fmt.Printf("Итоговый результат вычислений: %d\n", number)

	fmt.Println("Завершение программы.")
}
