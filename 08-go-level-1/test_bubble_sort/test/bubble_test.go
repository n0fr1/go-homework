package test

import (
	"testing"

	"github.com/n0fr1/go-homework/08-go-level-1/test_bubble_sort/bubblesort"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {

	in := []int{22, -1, 6, 9, 10, 0, 4, 1, -3, 8, 7, 13, 48} //входные данные

	out := []int{-3, -1, 0, 1, 4, 6, 7, 8, 9, 10, 13, 22, 48} //нужно получить

	result := bubblesort.Bubble(in) //результат сортировки

	for index := range out { //тестируем с использованием библиотеки testify
		t.Run("Testing", func(t *testing.T) {
			assert.Equal(t, out[index], result[index])
		})
	}

}
