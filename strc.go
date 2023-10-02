package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var oper_type string

func solution(text string) (string, string, string) {
	var first, sign, second string
	first, sign, second = edit(text)
	//Если вторая часть данных не строка то математическая операция
	//Если нет то операции со строками

	//secondArg, err := strconv.Atoi(second)
	if oper_type == "strings" {
		switch {
		case sign == "+":
			println(strings.ReplaceAll(first+second, "\"", ""))
		case sign == "-":
			resultf := strings.ReplaceAll(first, "\"", "")
			results := strings.ReplaceAll(second, "\"", "")
			if (len(resultf) > 10) || (len(results) > 10) == true {
				println("Нельзя вводить строки длиннее 10 символов!")
			}
			result := strings.ReplaceAll(resultf, results, "")
			println(result)
		default:
			panic("Строки между собой можно только складывать и вычитать, не знаю что с этим делать...")
		}
	} else {
		if sign == "*" {
			secondArg, _ := strconv.Atoi(second)
			if secondArg > 10 {
				panic("Нельзя вводить числа больше 10")
			}
			firstArg := strings.ReplaceAll(first, "\"", "")
			var printString string
			for i := 0; i < secondArg; i++ {
				//print(first)
				printString += firstArg
			}
			if len(printString) > 40 {
				println(printString[0:40], "...")
			} else {
				println(printString)
			}
		} else if sign == "/" {
			secnd, _ := strconv.Atoi(second)
			if secnd > 10 {
				panic("Нельзя вводить числа больше 10")
			}
			first = strings.ReplaceAll(first, "\"", "")
			fl := len(first)
			result := fl / secnd
			println(first[0:result])
		}

	}

	return first, second, sign
}

func edit(text string) (string, string, string) {
	var first, sign, second string
	var slice []string
	text = strings.TrimSpace(text)
	//text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	if strings.Contains(text, " + ") {
		slice = strings.Split(text, " + ")
		sign = "+"
	} else if strings.Contains(text, " - ") {
		slice = strings.Split(text, " - ")
		sign = "-"
	} else if strings.Contains(text, " * ") {
		slice = strings.Split(text, " * ")
		sign = "*"
	} else if strings.Contains(text, " / ") {
		slice = strings.Split(text, " / ")
		sign = "/"
	}
	first = slice[0]
	//sign = slice[1]
	second = slice[1]
	//Проверка на тип данных в первом(что подразумевалось string или int)
	if strings.Contains(first, "\"") == false {
		panic("Первым вводится строка в кавычках!!!")
	} else {
		//Если второе число содержит кавычки то
		if strings.Contains(second, "\"") {
			oper_type = "strings"
		} else {
			oper_type = "strint"
		}
	}

	return first, sign, second
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		println("\nВведите две строки и знак между ними.")
		text, _ := reader.ReadString('\n')
		solution(text)
	}

}

