package fileParser

import (
	"bufio"
	mower "exo/mower"
	"log"
	"os"
	"regexp"
	"strconv"
)

const gridRegexp = "^([0-9])([0-9])$"
const mowerRegexp = "^([0-9])([0-9]) ([NSWE])$"
const programRegexp = "^([ADG]*)$"

func ReadInputFile(filePath string) (x int, y int, mowers []mower.Mower){
	// Check if file exist
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scan line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid, _ := regexp.Compile(gridRegexp)
		mowerPosition, _ := regexp.Compile(mowerRegexp)
		programString, _ := regexp.Compile(programRegexp)

		// Extract grid info
		if grid.MatchString(scanner.Text()) {
			x, _ = strconv.Atoi(grid.FindStringSubmatch(scanner.Text())[1])
			y, _ = strconv.Atoi(grid.FindStringSubmatch(scanner.Text())[2])	
		} else if mowerPosition.MatchString(scanner.Text()) {
			// Extract mower init posistion
			program := ""
			
			mox, _ := strconv.Atoi(mowerPosition.FindStringSubmatch(scanner.Text())[1])
			moy, _ := strconv.Atoi(mowerPosition.FindStringSubmatch(scanner.Text())[2])
			direction := mowerPosition.FindStringSubmatch(scanner.Text())[3]
			// Read next line to get mower program
			if scanner.Scan() {
				if programString.MatchString(scanner.Text()) {
					program = programString.FindStringSubmatch(scanner.Text())[0]
				}
			}
			mo := mower.New(direction, mox, moy, program)
			mowers = append(mowers, mo)
			log.Printf("Mower Initialized : x:%d, y:%d, orientation:%s, prog:%s \n", mox, moy, direction, program)
		}

	}
	// Comfirm data are present
	log.Printf("Grid x:%d, y:%d", x, y)
	if x ==0 || y ==0 || len(mowers) == 0 {
		panic("Invalid input file")
	}
	return
}
