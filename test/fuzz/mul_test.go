package fuzz

import (
	"testing"
)

func FuzzMul(f *testing.F) {
	f.Fuzz(func(t *testing.T, start uint, end uint) {
		res , err:= Mul(start, end)
		if err != nil{
			t.Errorf("error occured result is %d ! %s",res, err.Error())
		}
	})
}
