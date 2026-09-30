package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2. Напиши функцию countLines(filename string) (int, error),
// которая читает файл построчно через bufio.Scanner и возвращает количество строк.
// Проверь на файле из задания 1.
func task2() {
	lines, err := countLines("text.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("lines:", lines)
}

func countLines(filename string) (int, error) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Ошибка закрытия файла:", err)
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения сканером:", err)
		return 0, err
	}

	return i, nil
}
