package main

// Пишите тесты в этом файле

import "testing"

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	// Тест с нормальным размером
	size := 10
	result := generateRandomElements(size)
	if len(result) != size {
		t.Error("Длина слайса не подходит")
	}

	// Тест с граничным условием - размер 0
	size = 0
	result = generateRandomElements(size)
	if result != nil {
		t.Error("При size=0 ожидается nil")
	}

	// Тест с отрицательным размером
	size = -5
	result = generateRandomElements(size)
	if result != nil {
		t.Error("При отрицательном size ожидается nil")
	}
}

func TestMaximum(t *testing.T) {
	// Тест с нормальными данными
	data := []int{3, 6, 4, 3, 1}
	slc := maximum(data)
	expected := 6 // максимальное значение в слайсе 6, а не 9
	if slc != expected {
		t.Errorf("Ошибка: ожидается %d, получено %d", expected, slc)
	}

	// Тест с одним элементом
	data = []int{42}
	slc = maximum(data)
	expected = 42
	if slc != expected {
		t.Errorf("Ошибка: ожидается %d, получено %d", expected, slc)
	}

	// Тест с отрицательными числами
	data = []int{-5, -2, -10, -1}
	slc = maximum(data)
	expected = -1
	if slc != expected {
		t.Errorf("Ошибка: ожидается %d, получено %d", expected, slc)
	}

	// Тест с пустым слайсом
	data = []int{}
	slc = maximum(data)
	expected = 0
	if slc != expected {
		t.Errorf("Ошибка: ожидается %d, получено %d", expected, slc)
	}
}
