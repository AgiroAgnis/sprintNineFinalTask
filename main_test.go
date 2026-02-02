package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	size := 10
	elements := generateRandomElements(size)
	if len(elements) != size {
		t.Errorf("Ожидаемая длина %d, реальная %d", size, len(elements))
	}

	for _, v := range elements {
		if v < 0 || v >= 10000 {
			t.Errorf("Элемент %d вне ожидаемого диапазона", v)
		}
	}
}

func TestMaximum(t *testing.T) {
	data := []int{1, 3, 2, 5, 4}
	expectedMax := 5
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{1, 3, 2, 5, 4, 8, 7, 6}
	expectedMax := 8
	actualMax := maxChunks(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}
}

func TestMaxChunksEmptySlice(t *testing.T) {
	data := []int{}
	expectedMax := 0
	actualMax := maxChunks(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	data := []int{}
	expectedMax := 0
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}
}

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	size := 0
	elements := generateRandomElements(size)
	if len(elements) != 0 {
		t.Errorf("Ожидаемая длина 0, реальная %d", len(elements))
	}
}

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	size := -1
	elements := generateRandomElements(size)
	if len(elements) != 0 {
		t.Errorf("Ожидаемая длина 0, реальная %d", len(elements))
	}
}

func TestOneElementSlice(t *testing.T) {
	data := []int{42}
	expectedMax := 42
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}

	actualMaxChunks := maxChunks(data)
	if actualMaxChunks != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMaxChunks)
	}
}

func TestAllElementsEqual(t *testing.T) {
	data := []int{7, 7, 7, 7, 7}
	expectedMax := 7
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}

	actualMaxChunks := maxChunks(data)
	if actualMaxChunks != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMaxChunks)
	}
}

func TestOnlyNegativeElements(t *testing.T) {
	data := []int{-10, -20, -5, -30}
	expectedMax := -5
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}

	actualMaxChunks := maxChunks(data)
	if actualMaxChunks != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMaxChunks)
	}
}

func TestMixedElements(t *testing.T) {
	data := []int{-10, 0, 5, -3, 8, 2}
	expectedMax := 8
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}

	actualMaxChunks := maxChunks(data)
	if actualMaxChunks != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMaxChunks)
	}
}

func TestSliceLessChunks(t *testing.T) {
	data := []int{3, 1}
	expectedMax := 3
	actualMax := maxChunks(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}
}

func TestSliceChunksSize(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedMax := 8
	actualMax := maxChunks(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}
}

func TestLargeSlice(t *testing.T) {
	size := 1000000
	data := generateRandomElements(size)
	maxValue := maximum(data)
	maxChunksValue := maxChunks(data)
	if maxValue != maxChunksValue {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", maxValue, maxChunksValue)
	}
}

func TestAllZeroes(t *testing.T) {
	data := []int{0, 0, 0, 0, 0}
	expectedMax := 0
	actualMax := maximum(data)
	if actualMax != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMax)
	}

	actualMaxChunks := maxChunks(data)
	if actualMaxChunks != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMaxChunks)
	}
}

func TestSliceNotDivisibleByChunks(t *testing.T) {
	data := []int{1, 5, 2, 9, 3}
	expectedMax := 9
	actualMaxChunks := maxChunks(data)
	if actualMaxChunks != expectedMax {
		t.Errorf("Ожидаемое максимальное значение %d, реальное %d", expectedMax, actualMaxChunks)
	}
}
