package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	// Табличные тесты
	tests := []struct {
		name     string
		size     int
		expected int  // ожидаемая длина
		wantNil  bool // ожидаем nil?
	}{
		{
			name:     "нормальный размер",
			size:     10,
			expected: 10,
			wantNil:  false,
		},
		{
			name:     "граничное условие - размер 0",
			size:     0,
			expected: 0,
			wantNil:  true,
		},
		{
			name:     "отрицательный размер",
			size:     -5,
			expected: 0,
			wantNil:  true,
		},
		{
			name:     "большой размер",
			size:     1000,
			expected: 1000,
			wantNil:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if tt.wantNil {
				assert.Nil(t, result, "ожидается nil")
			} else {
				assert.NotNil(t, result, "результат не должен быть nil")
				assert.Len(t, result, tt.expected, "длина слайса не соответствует")
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	// Табличные тесты
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "нормальные данные",
			data:     []int{3, 6, 4, 3, 1},
			expected: 6,
		},
		{
			name:     "один элемент",
			data:     []int{42},
			expected: 42,
		},
		{
			name:     "отрицательные числа",
			data:     []int{-5, -2, -10, -1},
			expected: -1,
		},
		{
			name:     "пустой слайс",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "все одинаковые числа",
			data:     []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "положительные и отрицательные",
			data:     []int{-10, 0, 10, -5, 5},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			assert.Equal(t, tt.expected, result, "максимальное значение не соответствует")
		})
	}
}

// Дополнительный тест для maxChunks
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "нормальные данные",
			data:     []int{3, 6, 4, 3, 1, 10, 2, 8},
			expected: 10,
		},
		{
			name:     "пустой слайс",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "один элемент",
			data:     []int{100},
			expected: 100,
		},
		{
			name:     "не делится на чанки",
			data:     []int{1, 5, 3, 9, 2, 8, 4, 7, 6}, // 9 элементов
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.data)
			assert.Equal(t, tt.expected, result, "максимальное значение не соответствует")
		})
	}
}
