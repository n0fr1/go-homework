package test

import (
	"reflect"
	"testProject/test_bubble_sort/bubblesort"
	"testing"
)

func TestBubble(t *testing.T) {

	in := []int{22, -1, 6, 9, 10, 0, 4, 1, -3, 8, 7, 13, 48} //входные данные

	out := []int{-3, -1, 0, 1, 4, 6, 7, 8, 9, 10, 13, 22, 48} //нужно получить

	result := bubblesort.Bubble(in) //результат

	for index := range out {

		t.Run("Testing", func(t *testing.T) {

			if !reflect.DeepEqual(out[index], result[index]) {
				t.Error("Expected ", out[index], " got ", result[index])
			}

		})

	}

}
