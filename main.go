package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Калькулятор сложения, деления, умножения, деления римских и арабских цифр \n")
	fmt.Println("Умеет работать с арабскими (от 1 до 10) или римскими (от I до X) \n")
	fmt.Print("Введите пример (в формате 10 * 5): ")
	text, _ := reader.ReadString('\n')

	arrayStrings := strings.Fields(text)
	if len(arrayStrings) != 3 {
		fmt.Println("Неверно введен пример")
	} else {
		one_value_str := arrayStrings[0]
		two_value_str := arrayStrings[2]
		operand_str := arrayStrings[1]

		one_number_str, indexArabianRoman1 := check_type(one_value_str)
		two_number_str, indexArabianRoman2 := check_type(two_value_str)

		o_c := operand_check(operand_str)

		if indexArabianRoman1 == 0 && indexArabianRoman2 == 0 {
			one_number_int, err := strconv.Atoi(one_number_str)

			if err != nil {
				panic(err)
			}
			two_number_int, err := strconv.Atoi(two_number_str)

			if err != nil {
				panic(err)
			}

			result := calculation(one_number_int, two_number_int, o_c)
			fmt.Println(result)

		} else if indexArabianRoman1 == 1 && indexArabianRoman2 == 1 {
			s1 := convert_arabic(one_number_str)
			s2 := convert_arabic(two_number_str)
			result := calculation(s1, s2, o_c)

			if result < 0 {
				panic("Ошибка: Итог уравнения римских цифр является отрциательное число")
			} else {
				result := convert_roman(result)
				fmt.Println(result)
			}
		} else {
			panic("Неверное значение")
		}
	}
}

func check_type(value string) (string, int) {
	arabic_number_list := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman_number_list := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	if slices.Contains(arabic_number_list, value) {
		return value, 0
	} else if slices.Contains(roman_number_list, value) {
		return value, 1
	} else {
		panic("Введеное неверное число. Числа должны быть арабскими (от 1 до 10) или римскими (от I до X)")
	}
}

func operand_check(operand string) string {
	if operand == "+" {
		return operand
	} else if operand == "-" {
		return operand
	} else if operand == "*" {
		return operand
	} else if operand == "/" {
		return operand
	} else {
		panic("Неверный операнд")
	}
}

func calculation(a, b int, operand string) int {
	if operand == "+" {
		return a + b
	} else if operand == "-" {
		return a - b
	} else if operand == "*" {
		return a * b
	} else if operand == "/" {
		if b == 0 {
			panic("На ноль делить нельзя")
		} else {
			return a / b
		}
	} else {
		return 0
	}
}

func convert_roman(n int) string {
	result := ""

	elements := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	keys := make([]int, len(elements))
	i := 0

	for k := range elements {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	slices.Reverse(keys)

	for _, p := range keys {
		y := n / p
		result += strings.Repeat(elements[p], y)
		n %= p

	}
	return result
}

func convert_arabic(n string) int {
	elements := map[string]int{
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	rome_numbers := n
	result := 0

	for n, i := range rome_numbers {
		for roman, arabian := range elements {
			if string(i) == roman {
				if n == len(rome_numbers)-1 {
					result += arabian
				} else {
					if string(i) == "C" && string(rome_numbers[n+1]) == "M" {
						result -= arabian
						continue
					}
					if string(i) == "C" && string(rome_numbers[n+1]) == "D" {
						result -= arabian
						continue
					}
					if string(i) == "X" && string(rome_numbers[n+1]) == "C" {
						result -= arabian
						continue
					}
					if string(i) == "X" && string(rome_numbers[n+1]) == "L" {
						result -= arabian
						continue
					}
					if string(i) == "I" && string(rome_numbers[n+1]) == "X" {
						result -= arabian
						continue
					}
					if string(i) == "I" && string(rome_numbers[n+1]) == "V" {
						result -= arabian
						continue
					} else {
						result += arabian
					}
				}
			}
		}
	}
	return result
}
