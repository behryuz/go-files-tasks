package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1. Сохрани список из 5 задач (строки) в файл tasks.txt через os.WriteFile,
// затем прочитай файл и выведи задачи в консоль пронумерованными, начиная с 1.

func task1() {
	fileContentMsg := "Задача 1\nЗадача 2\nЗадача 3\nЗадача 4\nЗадача 5"
	if err := os.WriteFile("test.txt", []byte(fileContentMsg), 0644); err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	fmt.Println("Файл записан")

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Ошибка закрытия файла:", err)
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		i++
		fmt.Printf("%d. %v\n", i, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения сканером:", err)
		return
	}
}
