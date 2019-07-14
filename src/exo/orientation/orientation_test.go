package orientation 

import (
	"testing"
)

func TestString2Orientation(t *testing.T) {
	if String2Orientation("N") != North() {
		t.Errorf ("Invalide conversion")
	}
	if String2Orientation("E") != East() {
		t.Errorf ("Invalide conversion")
	}
	if String2Orientation("S") != South() {
		t.Errorf ("Invalide conversion")
	}
	if String2Orientation("W") != West() {
		t.Errorf ("Invalide conversion")
	}
}	

func TestOrientation2String(t *testing.T) {
	if Orientation2String(North()) != "N" {
		t.Errorf ("Invalide conversion")
	}
	if Orientation2String(East()) != "E" {
		t.Errorf ("Invalide conversion")
	}
	if Orientation2String(South()) != "S" {
		t.Errorf ("Invalide conversion")
	}
	if Orientation2String(West()) != "W" {
		t.Errorf ("Invalide conversion")
	}
	if Orientation2String(-50) != "" {
		t.Errorf ("Invalide conversion")
	}
}	