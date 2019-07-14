package main

import (
	"log"
	"sync"
	"exo/parser"
	"os"
	controlcenter "exo/controlcenter"
	mower "exo/mower"
)

func main() {
	if len(os.Args) != 2 {
		panic("No Input file")	
	}
	log.Printf("Read Input File : %s", os.Args[1])

	gridX, gridY, mowers := fileParser.ReadInputFile(os.Args[1])
	var wg sync.WaitGroup
	ccc := make(chan mower.MoveRequest)
	go controlcenter.CheckMove(ccc, mowers, gridX, gridY)
	for _, mo := range mowers {
		wg.Add(1)
		go mo.Start(&wg, ccc)
	}

	wg.Wait()	
}

