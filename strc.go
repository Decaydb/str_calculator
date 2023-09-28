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
			//println(first + second)
		case sign == "-":
			result := strings.ReplaceAll(first, second, "")
			println(result)
		default:
			panic("Строки между собой можно только складывать и вычитать, не знаю что с этим делать...")
		}

		//Если первым введена не строка то ошибка
	} else {
		secondArg, _ := strconv.Atoi(second)
		for i := 0; i < secondArg; i++ {
			print(second)
		}
	}

	return first, second, sign
}

func edit(text string) (string, string, string) {
	var first, sign, second string
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	slice := strings.Split(text, " ")
	first = slice[0]
	sign = slice[1]
	second = slice[2]
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
		println("Введите две строки и знак между ними.")
		text, _ := reader.ReadString('\n')
		solution(text)
	}

}
