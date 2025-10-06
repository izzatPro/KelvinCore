package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	cMinimum = -273.15
	fMinimum = -459.67
	link     = "temperature.txt"
)

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5.0 / 9.0
}

func main() {
	fmt.Println("Добропожаловать в конвертор цельсия в фаренгейт !")
	fmt.Println("1. Цельсий в Фаренгейт")
	fmt.Println("2. Фаренгейт в Цельский")
	fmt.Print("Выберите тип конвертора (1 или 2): ")

	var choice int
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Println("Неверный символ вместо числа !")
		return
	}
	if choice != 1 && choice != 2 {
		fmt.Println("Введен не существующий номер выбора")
		return
	}

	switch choice {
	case 1:
		fmt.Print("Введите количество Цельсия которого нужно перевести в Фаренгейт: ")
		var celsius float64
		if _, err := fmt.Scanln(&celsius); err != nil {
			fmt.Println("Вы ввели неверное значение температуры")
			return
		}
		if celsius < cMinimum {
			fmt.Println("Вы ввели ниже минимального (-273.15) допустимого для цельсия ! ")
			return
		}
		fahrenheit := celsiusToFahrenheit(celsius)
		writeToFile(link, fahrenheit, "°F")
		fmt.Printf("%.2f Цельсия в Фаренгейтах: %.2f", celsius, fahrenheit)
	case 2:
		fmt.Print("Введите количество Фарнгейта которого нужно перевести в Цельсий: ")
		var fahrenheit float64
		if _, err := fmt.Scanln(&fahrenheit); err != nil {
			fmt.Println("Вы ввели неверное значение температуры в фарнгейтах")
			return
		}
		if fahrenheit < fMinimum {
			fmt.Println("Вы ввели ниже минимального (-459.67) допустимого для фаренгейта !")
			return
		}
		celsius := fahrenheitToCelsius(fahrenheit)
		writeToFile(link, celsius, "°C")
		fmt.Printf("%.2f Фаренгейта в Цельсиях: %.2f\n", fahrenheit, celsius)
		data, err := readFromFile()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
	default:
		return
	}
}

func writeToFile(link string, temperature float64, unit string) {
	temperatureToString := fmt.Sprintf("%.2f%s\n", temperature, unit)
	err := os.WriteFile(link, []byte(temperatureToString), 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}
}

func readFromFile() (string, error) {
	data, err := os.ReadFile(link)
	if err != nil {
		fmt.Println("Мы не можем прочитать файл с таким названием.")
		return "", errors.New("Файла с таким названием не существует")
	}
	dataToString := string(data)
	return dataToString, nil
}
