package test

import (
	"testProject/src/bubblesort"
	"testing"
)

func TestBubble(t *testing.T) {

	var want []int                                             //пока пробуем простой вариант теста
	need := []int{-3, -1, 0, 1, 4, 6, 7, 8, 9, 10, 13, 22, 48} //нужно получить

	vTest := []int{22, -1, 6, 9, 10, 0, 4, 1, -3, 8, 7, 13, 48} //входные данные

	want = bubblesort.Bubble(vTest) //результат

	for index := range want {

		if want[index] != need[index] {
			t.Error("Expected ", need[index], " got ", want[index])
		}

	}

}
