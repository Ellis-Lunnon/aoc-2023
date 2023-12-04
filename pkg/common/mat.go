package common

import "log"

type Mat[T any] struct {
	X, Y int
	Data []T
}

func (m *Mat[T]) Get(x int, y int) T {
	for x < 0 {
		x += m.X
	}
	for y < 0 {
		y += m.Y
	}
	return m.Data[m.X*x+y]
}

func (m *Mat[T]) GetXY(idx int) (int, int) {
	for idx < 0 {
		idx += len(m.Data)
	}
	return int(idx / m.Y), idx % m.Y
}

func (m *Mat[T]) GetIdx(x, y int) int {
	return m.X*x + y
}

func (m *Mat[T]) Set(x, y int, val T) {
	for x < 0 {
		x += m.X
	}
	for y < 0 {
		y += m.Y
	}
	m.Data[m.X*x+y] = val
}

func NewMat[T any](x, y int, content []T) Mat[T] {
	data := make([]T, x*y)
	if len(content) != x*y {
		log.Panicln("Incompatable shapes", x*y, len(content))
	}
	copy(data, content[:x*y])

	return Mat[T]{
		Data: data,
		X:    x,
		Y:    y,
	}
}
