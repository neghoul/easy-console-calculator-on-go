package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var a int
		var b int
		var exit string
		var choice string

		fmt.Print("Введите первое число: ")
		fmt.Scan(&a)

		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Scan(&choice)

		if choice != "+" && choice != "-" && choice != "*" && choice != "/" {
			fmt.Println("Некоректная операция!")
			break
		}

		fmt.Print("Введите второе число: ")
		fmt.Scan(&b)

		if choice == "+" {
			fmt.Println("Сумма:", a+b)
		} else if choice == "-" {
			fmt.Println("Вычитание:", a-b)
		} else if choice == "*" {
			fmt.Println("Умножение:", a*b)
		} else if choice == "/" {
			fmt.Println("Деление:", a/b)
		}

		fmt.Print("Вы хотите выйти? (Y/n): ")
		fmt.Scan(&exit)

		if strings.ToLower(exit) == "y" {
			fmt.Println("Досвидания!")
			break
		} else {
			continue
		}
	}
}
