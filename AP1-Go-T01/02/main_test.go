package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// Функция для подмены os.Stdin
func setStdin(input string) (func(), error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(w, input)
	if err != nil {
		return nil, err
	}
	w.Close()

	oldStdin := os.Stdin
	os.Stdin = r
	return func() { os.Stdin = oldStdin }, nil
}

// Тест для обычного поведения с K меньшим, чем количество уникальных слов
func TestNormalBehavior(t *testing.T) {
	input := "aa bb cc aa cc cc cc aa ab ac bb\n3\n"
	expectedOutput := "cc aa bb\n"

	// Подменяем os.Stdin
	cleanup, err := setStdin(input)
	if err != nil {
		t.Fatalf("Failed to set stdin: %v", err)
	}
	defer cleanup()

	// Перенаправляем стандартный вывод
	var buf strings.Builder
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Запускаем основную функцию
	main()

	// Читаем вывод
	w.Close()
	io.Copy(&buf, r)
	output := buf.String()

	// Проверяем вывод
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

// Тест для передачи пустого списка слов
func TestEmptyWordList(t *testing.T) {
	input := "\n3\n" // Пустая строка слов
	expectedOutput := "\n" // Ожидаемый вывод — пустая строка

	// Подменяем os.Stdin
	cleanup, err := setStdin(input)
	if err != nil {
		t.Fatalf("Failed to set stdin: %v", err)
	}
	defer cleanup()

	// Перенаправляем стандартный вывод
	var buf strings.Builder
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Запускаем основную функцию
	main()

	// Читаем вывод
	w.Close()
	io.Copy(&buf, r)
	output := buf.String()

	// Проверяем вывод
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

// Тест для передачи списка слов, где K больше, чем число уникальных слов
func TestKGreaterThanUniqueWords(t *testing.T) {
	input := "aa bb cc aa cc cc cc aa ab ac bb\n10\n" // K больше, чем уникальных слов
	expectedOutput := "cc aa bb ab ac\n" // Ожидаемый вывод — все уникальные слова

	// Подменяем os.Stdin
	cleanup, err := setStdin(input)
	if err != nil {
		t.Fatalf("Failed to set stdin: %v", err)
	}
	defer cleanup()

	// Перенаправляем стандартный вывод
	var buf strings.Builder
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Запускаем основную функцию
	main()

	// Читаем вывод
	w.Close()
	io.Copy(&buf, r)
	output := buf.String()

	// Проверяем вывод
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

// Тест для случая, когда слова имеют одинаковую частоту
func TestEqualFrequencyWords(t *testing.T) {
	input := "aa bb cc dd ee ff gg hh ii jj\n5\n" // Все слова встречаются по одному разу
	expectedOutput := "aa bb cc dd ee\n" // Ожидаемый вывод — первые K слов в лексикографическом порядке

	// Подменяем os.Stdin
	cleanup, err := setStdin(input)
	if err != nil {
		t.Fatalf("Failed to set stdin: %v", err)
	}
	defer cleanup()

	// Перенаправляем стандартный вывод
	var buf strings.Builder
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Запускаем основную функцию
	main()

	// Читаем вывод
	w.Close()
	io.Copy(&buf, r)
	output := buf.String()

	// Проверяем вывод
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}