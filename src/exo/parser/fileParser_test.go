package fileParser

import "testing"

func TestReadInputFile(t *testing.T) {

	x, y, mowers := ReadInputFile("../../ressource/in.dat")

	if x != 5 || y !=5 || len(mowers) != 2 {
		t.Errorf ("File parser result is invalid")
	} 


}