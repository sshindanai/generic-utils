package gutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct[T, RETURN any] struct {
	arr      []T
	callback func(T) RETURN
	expected []RETURN
}

func TestMap(t *testing.T) {
	testcases := []TestStruct[int, int]{
		{
			arr: []int{1, 2, 3},
			callback: func(i int) int {
				return i * 2
			},
			expected: []int{2, 4, 6},
		},
	}
	for _, tc := range testcases {
		actual := Map(tc.arr, tc.callback)
		assert.Equal(t, tc.expected, actual)
	}
}
