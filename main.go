package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Линейные задачи
func sumOfDigits() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	sum := 0
	strNum := strconv.Itoa(num)
	for _, digit := range strNum {
        n, _ := strconv.Atoi(string(digit))
        sum += n
    }

    fmt.Println("Сумма цифр:", sum)

}

func celsiusToFahrenheit() {
	var tempC float64
	fmt.Print("Введите температуру в градусах Цельсия: ")
	fmt.Scan(&tempC)
	fahrenheit := tempC*9/5 + 32
	fmt.Printf("%.2f градусов Цельсия = %.2f градусов Фаренгейта\n", tempC, fahrenheit)
}

func doubleArray() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Введите элементы массива: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := range arr {
		arr[i] *= 2
	}
	fmt.Println("Удвоенный массив:", arr)
}

func joinStrings() {
	var n int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&n)
	arr := make([]string, n)
	fmt.Println("Введите строки:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	result := strings.Join(arr, " ")
	fmt.Println("Объединенная строка:", result)
}

func distanceBetweenPoints() {
	var x1, y1, x2, y2 float64
	fmt.Print("Введите координаты первой точки (x1, y1): ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Введите координаты второй точки (x2, y2): ")
	fmt.Scan(&x2, &y2)
	distance := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

// Задачи с условным оператором
func isEven() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func isLeapYear() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Невисокосный")
	}
}

func maxOfThree() {
	var a, b, c int
	fmt.Print("Введите три числа: ")
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println("Наибольшее число:", a)
	} else if b > c {
		fmt.Println("Наибольшее число:", b)
	} else {
		fmt.Println("Наибольшее число:", c)
	}
}

func ageCategory() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	if age < 13 {
		fmt.Println("Ребенок")
	} else if age < 18 {
		fmt.Println("Подросток")
	} else if age < 60 {
		fmt.Println("Взрослый")
	} else {
		fmt.Println("Пожилой")
	}
}

func divisibleBy3And5() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

// Задачи на циклы
func factorial() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println("Факториал:", result)
}

func fibonacci() {
	var n int
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&n)
	fibs := make([]int, n)
	if n > 0 {
		fibs[0] = 0
	}
	if n > 1 {
		fibs[1] = 1
	}
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	fmt.Println("Числа Фибоначчи:", fibs)
}

func reverseArray() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	slices.Reverse(arr)

	fmt.Println("Перевернутый массив:", arr)
}

func primeNumbers() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	primes := []int{}
	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	fmt.Println("Простые числа до", n, ":", primes)
}

func sumArray() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Сумма чисел в массиве:", sum)
}

// Главная функция
func main() {
	fmt.Println("Выберите часть задания: ")
	fmt.Println("1 - Линейные задачи")
	fmt.Println("2 - Задачи с условным оператором")
	fmt.Println("3 - Задачи на циклы")

	var part, task int
	fmt.Scan(&part)

	switch part {
	case 1:
		fmt.Println("Выберите задачу:")
		fmt.Println("1 - Сумма цифр числа")
		fmt.Println("2 - Преобразование температуры")
		fmt.Println("3 - Удвоение каждого элемента массива")
		fmt.Println("4 - Объединение строк")
		fmt.Println("5 - Расчет расстояния между двумя точками")
		fmt.Scan(&task)
		switch task {
		case 1:
			sumOfDigits()
		case 2:
			celsiusToFahrenheit()
		case 3:
			doubleArray()
		case 4:
			joinStrings()
		case 5:
			distanceBetweenPoints()
		default:
			fmt.Println("Некорректный выбор.")
		}
	case 2:
		fmt.Println("Выберите задачу:")
		fmt.Println("1 - Проверка на четность/нечетность")
		fmt.Println("2 - Проверка високосного года")
		fmt.Println("3 - Определение наибольшего из трех чисел")
		fmt.Println("4 - Категория возраста")
		fmt.Println("5 - Проверка делимости на 3 и 5")
		fmt.Scan(&task)
		switch task {
		case 1:
			isEven()
		case 2:
			isLeapYear()
		case 3:
			maxOfThree()
		case 4:
			ageCategory()
		case 5:
			divisibleBy3And5()
		default:
			fmt.Println("Некорректный выбор.")
		}
	case 3:
		fmt.Println("Выберите задачу:")
		fmt.Println("1 - Факториал числа")
		fmt.Println("2 - Числа Фибоначчи")
		fmt.Println("3 - Реверс массива")
		fmt.Println("4 - Поиск простых чисел")
		fmt.Println("5 - Сумма чисел в массиве")
		fmt.Scan(&task)
		switch task {
		case 1:
			factorial()
		case 2:
			fibonacci()
		case 3:
			reverseArray()
		case 4:
			primeNumbers()
		case 5:
			sumArray()
		default:
			fmt.Println("Некорректный выбор.")
		}

	default:
		fmt.Println("Некорректный выбор.")
	}
}
