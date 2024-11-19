package basket

import (
	asrt "github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate_Add(t *testing.T) {
	x, y := 3, 3

	c := Calculate{}

	actual := c.Add(x, y)
	expected := 6
	if actual != expected {
		t.Errorf("Calculate.Add(%d,%d) failed. Expected: %d, Actual:%d", x, y, expected, actual)
	}
}

func TestCalculate_Subtract(t *testing.T) {
	c := Calculate{}

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 0},
		{5, 3, 2},
		{5, 3, 2},
	}

	for _, v := range tables {
		actual := c.Subtract(v.x, v.y)
		if actual != v.expected {
			t.Errorf("Calculate.Add(%d,%d) failed. Expected: %d, Actual:%d", v.x, v.y,
				v.expected, actual)
		}
	}
}

func TestCalculate_Multiply(t *testing.T) {

	assert := asrt.New(t)
	c := Calculate{}

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 4},
		{5, 3, 15},
		{5, 3, 15},
	}

	for _, v := range tables {
		actual := c.Multiply(v.x, v.y)
		assert.Equal(v.expected, actual)
	}
}

func TestDivide(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}

	tables := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{2, 2, 1},
		{6, 2, 3},
	}

	for _, v := range tables {
		actual := c.Divide(v.x, v.y)

		assert.Equal(v.expected, actual)
	}
}
