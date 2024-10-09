package fuzz

import (
	"errors"
)

func Mul(start uint, end uint) (mul uint, err error) {

	mul = 1

	var i uint
	for i = start; i < end; i++ {
		mul *= 2
		if mul >= 1000 {
			err = errors.New("result is so long !\n")
			break
		}
	}

	return
}
