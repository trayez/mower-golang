package controlcenter

import (
	mower "exo/mower"
	"log"
)

func CheckMove(cc chan mower.MoveRequest, mowers []mower.Mower, x int, y int) {
	
	for {
		request, _ := <- cc
		log.Printf("CC : Request from %s recieved", request.Uuid.String())
		// Check if mower still on grid 
		if request.X > x  || request.X < 0 || request.Y > y  || request.Y < 0  {
			log.Printf("CC : out of range Move From %s", request.Uuid.String())
			panic("Stop all mowers")
		}
		// Check if mower next pos is available
		for _, mo := range mowers {
			if !mo.Is(request.Uuid) {
				if mo.IsOnPos(request.X, request.Y) {
					log.Printf("CC : Request from %s REFUSED", request.Uuid.String())
					log.Printf("CC : Colision Move From %s", request.Uuid.String())
					panic("Stop all mowers")
				}
			}
		}
		// Send move confirmation
		// TODO : Use Map instead of list
		for _, mo := range mowers {
			if mo.Is(request.Uuid) {
				log.Printf("CC : Request from %s AUTHORIED", request.Uuid.String())
				mo.Com <- mower.AUTHORIZED
			}
		}
	}
}