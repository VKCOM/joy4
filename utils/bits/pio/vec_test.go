package pio

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func ExampleVec() {
	vec := [][]byte{[]byte{1, 2, 3}, []byte{4, 5, 6, 7, 8, 9}, []byte{10, 11, 12, 13}}
	println(VecLen(vec))

	vec = VecSlice(vec, 1, -1)
	logrus.Info(vec)

	vec = VecSlice(vec, 2, -1)
	logrus.Info(vec)

	vec = VecSlice(vec, 8, 8)
	logrus.Info(vec)

	// Output:
}
