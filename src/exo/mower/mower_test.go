package mower

import (
	"testing"
	"exo/orientation"
	"time"

)

var mo =  New("E", 0, 0, "AAAAAAA")

func TestLeft(t *testing.T) {
	mo.left();
	if mo.orientation != orientation.North() {
		t.Errorf ("Invalide orientation")
	}
}	

func TestRight (t *testing.T) {
	mo.right();
	if mo.orientation != orientation.East() {
		t.Errorf ("Invalide orientation")
	}
}	

func TestIsOnPos (t *testing.T) {
	if !mo.IsOnPos(0, 0) {
		t.Errorf ("Invalide Position ")
	}
}			

func TestMove (t *testing.T) {
	cc := make(chan MoveRequest)
	go mo.move(cc);
	<- cc
	mo.Com <- AUTHORIZED
	time.Sleep(100 * time.Millisecond)
	if !mo.IsOnPos(1, 0) {
		t.Errorf ("Invalide orientation ")
	}
}	