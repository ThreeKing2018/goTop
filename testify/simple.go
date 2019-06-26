package main

import (
	"github.com/pkg/errors"
)

func main() {

}
func Add(a, b int) int {
	return a + b
}
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return a / b, nil
}