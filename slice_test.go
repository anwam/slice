package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	ast := assert.New(t)
	expected := []int{2, 4, 6, 8, 10}
	result := Map([]int{1, 2, 3, 4, 5}, func(i int) int {
		i = i * 2
		return i
	})
	ast.Equal(expected, result)
}

func Test_MapIntToString(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"1", "5", "10"}
	result := Map([]int{1, 5, 10}, func(i int) string {
		return strconv.Itoa(i)
	})
	assert.Equal(expected, result)
}
