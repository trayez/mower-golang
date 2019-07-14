package mower

import (
	"log"
	"math"
	"sync"
//	"time"
	"strings"
	"strconv"
	orientation "exo/orientation"
	"github.com/google/uuid"
)

const AUTHORIZED = "AUTORIZED"
const REFUSED = "REFUSED"

type Mower struct {
	uuid        uuid.UUID
	orientation orientation.Orientation
	posX        int
	posY        int
	program     string
	Com			chan string
}

type MoveRequest struct {
	Uuid        uuid.UUID
	X int
	Y int
}

func (m *Mower) right() {
	m.orientation = m.orientation + orientation.RotationIncrement
	log.Printf("Mower New orientation %s", m.String())
}

func (m *Mower) left() {
	m.orientation = m.orientation - orientation.RotationIncrement
	log.Printf("Mower New orientation %s", m.String())
}

func (m *Mower) move(cc chan MoveRequest) {
	// Calculate new potential pos
	x := m.posX + int(math.Round(math.Sin(float64(m.orientation) * (math.Pi / 180))))
	y := m.posY + int(math.Round(math.Cos(float64(m.orientation) * (math.Pi / 180))))
	
	// Send an authorization request to control center
	request := MoveRequest{m.uuid, x, y}
	log.Printf("Mower %s: Send request",  m.uuid)
	cc <- request

	// Wait for authorization
	authorization, _ := <- m.Com 
	log.Printf("Mower %s : response  %s", m.uuid, authorization)
	if authorization == AUTHORIZED {
		m.posX = x
		m.posY = y
		log.Printf(m.String())
	} 
}

func (m Mower) Start(wg *sync.WaitGroup, cc chan MoveRequest) {
	defer wg.Done()
	for _, action := range m.program {
		//s = s + (10 * time.Millisecond)
		switch string(action) {
		case "A":
			m.move(cc)
		case "D":
			m.right()
		case "G":
			m.left()
		default:
			log.Printf("Unknown action : %c", action)
			panic("Invalid action")
		}
		//time.Sleep(100 * time.Millisecond)
	}
	log.Printf("### Final Position %s", m.String())
}

func New(direction string, posX int, posY int, program string) (mo Mower) {
	mo.uuid, _ = uuid.NewRandom()
	mo.posX = posX
	mo.posY = posY
	mo.program = program
	mo.orientation = orientation.String2Orientation(direction)
	mo.Com = make(chan string)
	return
}

func (m *Mower) IsOnPos(x int, y int) (bool){
	return	m.posX == x && m.posY == y
}

func (m *Mower) Is(uuid uuid.UUID) (bool){
	return	uuid == m.uuid
}

func (m *Mower) String() (string){
	var sb strings.Builder
	sb.WriteString("Mower:")
	sb.WriteString(m.uuid.String())
	sb.WriteString(", Pos x:")
	sb.WriteString(strconv.Itoa(m.posX))
	sb.WriteString(", y:")
	sb.WriteString(strconv.Itoa(m.posY))
	sb.WriteString(", orientation:")
	sb.WriteString(orientation.Orientation2String(m.orientation))
	return sb.String()
  }


